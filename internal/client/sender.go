package client

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_chat/pkg/protocol"
)

func (c *Client) Send(transmission json.RawMessage) error {
	length := uint32(len(transmission))
	err := binary.Write(c.conn, binary.BigEndian, length)
	if err != nil {
		return fmt.Errorf("error writing length: %w", err)
	}

	_, err = c.conn.Write(transmission)
	if err != nil {
		return fmt.Errorf("error writing transmission: %w", err)
	}

	return nil
}

func (c *Client) SendAndAwait(uid uint32, transmission json.RawMessage) (protocol.Transmission, error) {
	responseChan := make(chan protocol.Transmission, 1)

	c.RegisterAwaiter(uid, responseChan)

	err := c.Send(transmission)
	if err != nil {
		c.awaitLock.Lock()
		delete(c.awaiters, uid)
		c.awaitLock.Unlock()
		return protocol.Transmission{}, err
	}

	response := <-responseChan
	return response, nil
}

func (c *Client) SendMessage(body string, sessionID int) error {
	msg := protocol.Message{ // ← Changed
		OriginId:  c.GetID(),
		Body:      body,
		SessionId: sessionID,
	}

	_, transmission, err := protocol.NewTransmission(protocol.MsgCode, msg) // ← Changed
	if err != nil {
		return fmt.Errorf("error creating message: %w", err)
	}

	return c.Send(transmission)
}

func (c *Client) SendConnectionRequest(targetID int) (protocol.Transmission, error) {
	request := protocol.ConnectionRequest{ // ← Changed
		Requester: c.GetID(),
		Target:    targetID,
	}

	uid, transmission, err := protocol.NewTransmission(protocol.ConnectionRequestCode, request) // ← Changed
	if err != nil {
		return protocol.Transmission{}, fmt.Errorf("error creating request: %w", err)
	}

	response, err := c.SendAndAwait(uid, transmission)
	if err != nil {
		return protocol.Transmission{}, fmt.Errorf("error sending request: %w", err)
	}

	return response, nil
}

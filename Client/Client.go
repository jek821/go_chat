package main

import (
	"encoding/json"
	"fmt"
	"go_chat/Protocol"
	"go_chat/Shared"
	"go_chat/Utils"
	"net"
)

const serverIp = "127.0.0.1:8080"

type Client struct {
	conn net.Conn
	ID   int
	Shared.ListenerLogic
	Shared.SenderLogic
	Awaiters Shared.AwaitMap
}

func NewClient() *Client {

	// connect
	conn, err := net.Dial("tcp", serverIp)
	if err != nil {
		Utils.HandleErr(err)
	}

	// Set
	clientId := -1
	// store the client
	client := &Client{
		conn:     conn,
		ID:       clientId,
		Awaiters: Shared.CreateAwaitMap(),
	}

	return client
}

func (c *Client) HandleDisconnect() {
	fmt.Println("Connection To Server Lost")
	err := c.conn.Close()
	if err != nil {
		Utils.HandleErr(err)
	}
	return

}

func (c *Client) RunListener() {
	c.HandleIncomingPayLoads(c.conn, c.PayloadHandler, c.HandleDisconnect)
}

func (c *Client) PayloadHandler(p Protocol.Payload) {
	// First Check Awaiters
	_, ok := c.Awaiters.Map[p.Pid]
	if ok {
		c.Awaiters.ResolveWaiter(p)
		return
	}

	switch p.Code {
	case Protocol.GiveClientIdCode:
		var data Protocol.GiveClientId
		err := json.Unmarshal(p.Data, &data)
		if err != nil {
			Utils.HandleErr(err)
		}
		fmt.Printf("Received ID: %d\n", data.Id)

	default:
		fmt.Println("Unknown Payload Code")
	}

}

func (c *Client) SendTestMessage() {

}

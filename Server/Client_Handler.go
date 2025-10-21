package main

import (
	"encoding/json"
	"fmt"
	"go_chat/Protocol"
	"go_chat/Shared"
	"go_chat/Utils"
	"net"
)

type ClientHandler struct {
	conn  net.Conn
	CliId int
	Shared.ListenerLogic
	Shared.SenderLogic
	ServChan chan Protocol.Payload
}

func NewCliHandler(conn net.Conn, CliId int, ServChan chan Protocol.Payload) *ClientHandler {
	fmt.Println("New Client Handler Spawned")
	cliHandler := ClientHandler{conn: conn, CliId: CliId, ServChan: ServChan}
	cliHandler.runListener()
	return &cliHandler
}

func (c *ClientHandler) HandleDisconnect() {
	fmt.Printf("Connection To Client %d Lost\n", c.CliId)
	err := c.conn.Close()
	data := Protocol.EndClient{Id: c.CliId}
	encodedData, err := json.Marshal(data)
	if err != nil {
		Utils.HandleErr(err)
	}
	c.ServChan <- Protocol.Payload{Code: Protocol.EndClientCode, Data: encodedData}
	if err != nil {
		Utils.HandleErr(err)
	}
	return

}

func (c *ClientHandler) runListener() {
	c.GiveClientId()
	go c.HandleIncomingPayLoads(c.conn, c.payloadHandler, c.HandleDisconnect)
}

func (c *ClientHandler) payloadHandler(payload Protocol.Payload) {
	switch payload.Code {

	default:
		fmt.Println("Unrecognized Code in Payload")
	}

}

func (c *ClientHandler) GiveClientId() {
	var payload Protocol.Payload
	content := Protocol.GiveClientId{Id: c.CliId}
	encodedContent, err := json.Marshal(content)
	if err != nil {
		Utils.HandleErr(err)
	}
	payload.Code = Protocol.GiveClientIdCode
	payload.Data = encodedContent
	c.SendPayload(c.conn, payload)
}

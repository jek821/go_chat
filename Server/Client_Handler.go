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
}

func NewCliHandler(conn net.Conn, CliId int) *ClientHandler {
	fmt.Println("New Client Handler Spawned")
	cliHandler := ClientHandler{conn: conn, CliId: CliId}
	return &cliHandler
}

func (c *ClientHandler) runListener() {
	go c.HandleIncomingPayLoads(c.conn, c.payloadHandler)
}

func (c *ClientHandler) payloadHandler(payload Protocol.Payload) {
	switch payload.Code {
	case Protocol.RequestClientIdCode:
		content := Protocol.GiveClientId{Id: c.CliId}
		encodedContent, err := json.Marshal(content)
		if err != nil {
			Utils.HandleErr(err)
		}
		payload.Code = Protocol.GiveClientIdCode
		payload.Data = encodedContent
		c.SendPayload(c.conn, payload)
	default:
		fmt.Println("Unrecognized Code in Payload")
	}

}

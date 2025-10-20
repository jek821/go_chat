package main

import (
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

func (c *Client) GetID() int {
	// Sends a request to the server for a new client ID
	return c.ID
}

func (c *Client) RunListener() {
	go c.HandleIncomingPayLoads(c.conn, c.PayloadHandler)
}

func (c *Client) PayloadHandler(p Protocol.Payload) {
	// First Check Awaiters
	_, ok := c.Awaiters.Map[p.Pid]
	if ok {
		c.Awaiters.ResolveWaiter(p)
	}

	switch p.Code {
	case Protocol.TestCode:
	default:
		fmt.Println("Unknown Payload Code")
	}

}

func (c *Client) SendTestMessage() {

}

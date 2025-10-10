package main

import (
	"fmt"
	"go_chat/utils"
	"net"
)

type ClientHandler struct {
	id        int
	conn      net.Conn
	channelIn chan utils.Transmission
}

func clientReader(cli *ClientHandler) (utils.Transmission, error) {
	buffer := make([]byte, 1024)
	numBytes, err := cli.conn.Read(buffer)
	if err != nil {
		return utils.Transmission{}, err
	}
	encodedMsg := buffer[0:numBytes]
	transmission, err := utils.UnmarshalData(encodedMsg)
	if err != nil {
		return utils.Transmission{}, err
	}
	return transmission, nil
}

func cliHandlerFactory(id int, conn net.Conn, serverChan chan utils.Transmission) ClientHandler {
	newHandler := ClientHandler{
		id:        id,
		conn:      conn,
		channelIn: serverChan,
	}
	return newHandler
}

func giveCliNewId(client *ClientHandler) error {
	var code utils.Code = utils.GiveClientNewId
	var newId = utils.GiveClientId{Code: code, Id: client.id}
	var trans, err = utils.TransmissionFactory(code, newId, client.id)
	if err != nil {
		return err
	}
	writeToClient(client, trans)
	return nil
}
func cliHandler(client *ClientHandler) error {
	defer client.conn.Close()
	fmt.Printf("Client Handler %d started\n", client.id)
	giveCliNewId(client)

	for {
		data, err := clientReader(client)
		if err != nil {
			fmt.Printf("Client %d disconnected: %v\n", client.id, err)
			return err
		}

		// fmt.Printf("Client Handler %d Received Transmission: Code=%d, Data=%v\n",
		// 	client.id, data.Code, data.Data)

		// Route the transmission to the server channel for processing
		if client.channelIn != nil {
			client.channelIn <- data
		}
	}
}

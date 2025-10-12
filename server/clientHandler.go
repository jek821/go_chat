package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_chat/utils"
	"net"
)

type ClientHandler struct {
	id        int
	conn      net.Conn
	channelIn chan utils.Transmission
}

// Main Function For Client Handler
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

func clientReader(cli *ClientHandler) (utils.Transmission, error) {
	var length uint32
	binary.Read(cli.conn, binary.BigEndian, &length)
	buffer := make([]byte, length)
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
	var code utils.Code = utils.GiveClientNewIdCode
	var newId = utils.GiveClientId{Id: client.id}
	var _, trans, err = utils.TransmissionFactory(code, newId)
	if err != nil {
		return err
	}
	writeToClient(client, trans)
	return nil
}

func writeToClient(client *ClientHandler, transmission json.RawMessage) error {
	_, err := client.conn.Write(transmission)
	if err != nil {
		return err
	}
	return nil

}

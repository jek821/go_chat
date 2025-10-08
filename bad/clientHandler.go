package main

import (
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

func unmarshalData(msg []byte) utils.Transmission {
	var decodedTransmission utils.Transmission
	err := json.Unmarshal(msg, &decodedTransmission)
	if err != nil {
		fmt.Println("Error Decoding Message")
	}
	return decodedTransmission

}

func clientReader(cli *ClientHandler) (utils.Transmission, error) {
	buffer := make([]byte, 1024)
	numBytes, err := cli.conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading connection into buffer or Client Disconnected")
		return utils.Transmission{}, err
	}
	encodedMsg := buffer[0:numBytes]
	var transmission = unmarshalData(encodedMsg)
	return transmission, nil
}

func cliHandlerFactory(id int, conn net.Conn, serverChan chan utils.Transmission) ClientHandler {
	var newHandler ClientHandler = ClientHandler{id: id, conn: conn, channelIn: serverChan}
	return newHandler
}

func cliHandler(client *ClientHandler) error {
	defer client.conn.Close()
	for {
		data, err := clientReader(client)
		var msg = data.Data
		if err != nil {
			return err
		}
		fmt.Printf("New Client ID: %d\n", client.id)
		fmt.Printf("Client Handler %d Received Transmission: %s", client.id, msg)
	}
}

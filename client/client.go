package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"go_chat/utils"
	"net"
	"os"
)

var ClientId int

func main() {
	conn, err := startCli()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	// Start message sending loop
	go handleServer(conn)
	sendMsg(conn)
}

func startCli() (net.Conn, error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return nil, errors.New("ERROR CONNECTING TO SERVER")
	}

	fmt.Println("Connection To Server Successful!")
	return conn, nil
}

func sendMsg(conn net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("input text:")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error Reading Input")
			continue
		}

		msg := utils.Message{Body: text}

		// Marshal the message into JSON bytes first
		msgBytes, err := json.Marshal(msg)
		if err != nil {
			fmt.Println("Error encoding message:", err)
			continue
		}

		// Now assign the bytes to Data
		transmission := utils.Transmission{Code: utils.Msg, Data: msgBytes}
		data := encodeData(transmission)

		_, err = conn.Write(data)
		if err != nil {
			fmt.Println("Error Writing Message to Conn")
			return
		}
	}
}

func encodeData(data utils.Transmission) []byte {
	EncodedTransaction, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error Encoding Message: %v\n", err)
	}
	return EncodedTransaction
}

// func generateConnectionReq(requester int, target int) utils.ConnectionRequest {
// 	connReq := utils.ConnectionRequest{Requester: requester, Target: target}
// 	return connReq
// }

func serverReader(conn net.Conn) (utils.Transmission, error) {
	buffer := make([]byte, 1024)
	numBytes, err := conn.Read(buffer)
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

func handleServer(conn net.Conn) error {
	defer conn.Close()
	for {
		data, err := serverReader(conn)
		if err != nil {
			return err
		}
		switch data.Code {
		case utils.GetId:
			var getId utils.GiveClientId

			err := json.Unmarshal(data.Data, &getId)
			if err != nil {
				fmt.Println("Error unmarshaling Message:", err)
				continue
			}
			fmt.Printf("Client ID Received: %d\n", getId.Id)
		default:
			fmt.Printf("Unknown message code: %d\n", data.Code)

		}

	}

}

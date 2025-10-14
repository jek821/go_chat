package main

import (
	"encoding/json"
	"fmt"
	"go_chat/utils"
	"net"
)

var Port = 8080
var Ip = "127.0.0.1"
var clientPipe chan utils.Transmission
var clientIdCount = 0
var clients = make(map[int]*ClientHandler)
var sessions = make(map[int]*utils.Session)

func main() {
	// Initialize the channel
	clientPipe = make(chan utils.Transmission, 100)

	l := runServ()
	fmt.Printf("Server listening on %s:%d\n", Ip, Port)

	// Start a goroutine to process messages from the channel
	go processMessages()

	newConnHandler(l, clientPipe)
}

func runServ() net.Listener {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", Ip, Port))
	if err != nil {
		fmt.Println("Error starting server:", err)
		panic(err)
	}
	return l
}

func newConnHandler(l net.Listener, serverChan chan utils.Transmission) {
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		newID := generateUID()
		newCliHandler := cliHandlerFactory(newID, conn, serverChan)
		clients[newID] = &newCliHandler
		fmt.Printf("New client connected with ID: %d\n", newID)

		go cliHandler(&newCliHandler)
	}
}

func generateUID() int {
	clientIdCount++
	return clientIdCount
}

func processMessages() error {
	for {
		for trans := range clientPipe {
			fmt.Printf("Processing Transmission: Code=%d\n", trans.Code)
			switch trans.Code {
			case utils.MsgCode:
				var message utils.Message
				// The Unmarshalling done in the client handler results in a Transmission Struct
				// The Transmission struct's data is still in json byte code to allow for the abstraction of the contained struct type
				// Here we unmarshal the data into the correct struct (in this case message)
				err := json.Unmarshal(trans.Data, &message)
				if err != nil {
					return err
				}
				fmt.Printf("Message Received: %s\nFrom Client: %d\n", message.Body, message.OriginId)
			case utils.ConnectionRequestCode:
				var request utils.ConnectionRequest
				err := json.Unmarshal(trans.Data, &request)
				if err != nil {
					return err
				}
				_, ok := clients[request.Target]
				if ok {
					trans, err := json.Marshal(trans)
					if err != nil {
						return err
					}
					writeToClient(clients[request.Target], trans)
				}
			default:
				fmt.Printf("Unknown message code: %d\n", trans.Code)
			}
		}
	}
}

package main

import (
	"fmt"
	"go_chat/utils"
	"net"
)

var Port = 8080
var Ip = "127.0.0.1"
var clientPipe chan utils.Transmission
var clientIdCount = 0
var clients = make(map[int]*ClientHandler)

func main() {
	l := runServ()
	newConnHandler(l, clientPipe)
}

func runServ() net.Listener {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", Ip, Port))
	if err != nil {
		fmt.Println(err)
	}
	return l
}

func newConnHandler(l net.Listener, serverChan chan utils.Transmission) {
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		var newID = generateUID()
		newCliHandler := cliHandlerFactory(newID, conn, serverChan)
		clients[newID] = &newCliHandler
		go cliHandler(&newCliHandler)
	}

}

func generateUID() int {
	clientIdCount++
	return clientIdCount
}

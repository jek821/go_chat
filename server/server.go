package main

import (
	"encoding/json"
	"fmt"
	"net"
)

var Port = 8080
var Ip = "127.0.0.1"

func main() {
	runServ()
}

func runServ() {
	l, err := net.Listen("tcp", fmt.Sprintf("%s:%d", Ip, Port))
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	buffer := make([]byte, 1024)
	for {
		numBytes, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading connection into buffer or Client Disconnected")
			return
		}
		encodedMsg := buffer[0:numBytes]
		var decodedMsg = decoder(encodedMsg)
		fmt.Printf("Message Recieved: %s", decodedMsg)
	}
}

func decoder(msg []byte) string {
	var decodedMsg string
	err := json.Unmarshal(msg, &decodedMsg)
	if err != nil {
		fmt.Println("Error Decoding Message")
	}
	return decodedMsg

}

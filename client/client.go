package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func main() {
	runCli()
}

func runCli() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error connecting to server")
		return
	}

	defer conn.Close()

	fmt.Println("Connection To Server Successful!")

	takeMsg(conn)
}

func takeMsg(conn net.Conn) {
	for {
		fmt.Println("input text:")
		reader := bufio.NewReader(os.Stdin)
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error Reading Input")
		}
		data := encode(msg)
		conn.Write(data)
	}
}

func encode(msg string) []byte {
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Printf("Error Encoding Message: %s", msg)
	}
	return data

}

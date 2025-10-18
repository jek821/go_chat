package main

import (
	"fmt"
	"go_chat/internal/server"
)

func main() {
	srv := server.NewServer("10.1.1.2", 8080)

	fmt.Printf("Server listening on %s:%d\n", srv.Ip, srv.Port)

	if err := srv.Start(); err != nil {
		panic(err)
	}
}

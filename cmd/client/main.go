package main

import (
	"fmt"
	"go_chat/internal/client"
	"log"
)

func main() {
	fmt.Println("Starting chat client...")

	// Create and connect a client
	cli, err := client.NewClient(":8080")
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer cli.Close()

	// Start listener in the background
	go func() {
		if err := cli.Listener(); err != nil {
			log.Printf("Listener stopped: %v", err)
		}
	}()

	// Run the interactive menu (blocks until user quits)
	menu := NewMenu(cli)
	if err := menu.Run(); err != nil {
		log.Fatalf("Menu error: %v", err)
	}
}

package main

import (
	"bufio"
	"fmt"
	"go_chat/internal/client"
	"go_chat/pkg/protocol"
	"os"
	"strconv"
	"strings"
)

type Menu struct {
	client *client.Client
	reader *bufio.Reader
}

func NewMenu(cli *client.Client) *Menu {
	return &Menu{
		client: cli,
		reader: bufio.NewReader(os.Stdin),
	}
}

func (m *Menu) Run() error {
	for {
		m.displayMainMenu()

		choice, err := m.readInput()
		if err != nil {
			return err
		}

		switch strings.ToLower(choice) {
		case "1", "connect":
			if err := m.handleConnectionRequest(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		case "2", "message":
			if err := m.handleSendMessage(); err != nil {
				fmt.Printf("Error: %v\n", err)
			}
		case "3", "quit", "q":
			fmt.Println("Goodbye!")
			return nil
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func (m *Menu) displayMainMenu() {
	fmt.Println("\n" + strings.Repeat("=", 50))
	fmt.Printf("Chat Client (ID: %d)\n", m.client.GetID())
	fmt.Println(strings.Repeat("=", 50))
	fmt.Println("1. Connect to another client")
	fmt.Println("2. Send message")
	fmt.Println("3. Quit")
	fmt.Print("\nEnter your choice: ")
}

func (m *Menu) handleConnectionRequest() error {
	fmt.Print("Enter target client ID: ")
	targetID, err := m.readInt()
	if err != nil {
		return fmt.Errorf("invalid client ID: %w", err)
	}

	if targetID == m.client.GetID() {
		return fmt.Errorf("cannot connect to yourself")
	}

	fmt.Printf("Sending connection request to client %d...\n", targetID)

	response, err := m.client.SendConnectionRequest(targetID)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}

	switch response.Code {
	case protocol.SessionRequestAcceptedCode:
		fmt.Println("✓ Connection accepted!")
		return nil
	case protocol.SessionRequestRejectedCode:
		fmt.Println("✗ Connection rejected")
		return nil
	default:
		return fmt.Errorf("unexpected response code: %d", response.Code)
	}
}

func (m *Menu) handleSendMessage() error {
	fmt.Print("Enter session ID: ")
	sessionID, err := m.readInt()
	if err != nil {
		return fmt.Errorf("invalid session ID: %w", err)
	}

	fmt.Print("Enter message: ")
	message, err := m.readInput()
	if err != nil {
		return err
	}

	if strings.TrimSpace(message) == "" {
		return fmt.Errorf("message cannot be empty")
	}

	err = m.client.SendMessage(message, sessionID)
	if err != nil {
		return fmt.Errorf("failed to send message: %w", err)
	}

	fmt.Println("✓ Message sent!")
	return nil
}

func (m *Menu) readInput() (string, error) {
	text, err := m.reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func (m *Menu) readInt() (int, error) {
	text, err := m.readInput()
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(text)
}

package server

import (
	"fmt"
	"go_chat/internal/client"
	"go_chat/pkg/protocol"
	"net"
)

type ClientHandler struct {
	id        int
	transport client.Transport
	channelIn chan protocol.Transmission
}

func (s *Server) cliHandlerFactory(id int, conn net.Conn) ClientHandler {
	return ClientHandler{
		id:        id,
		transport: &client.TCPTransport{Conn: conn},
		channelIn: s.clientPipe,
	}
}

func (s *Server) cliHandler(client *ClientHandler) error {
	defer client.transport.Close()
	fmt.Printf("Client Handler %d started\n", client.id)

	s.giveCliNewId(client)

	for {
		data, err := s.clientReader(client)
		if err != nil {
			fmt.Printf("Client %d disconnected: %v\n", client.id, err)
			delete(s.clients, client.id)
			return err
		}

		if client.channelIn != nil {
			client.channelIn <- data
		}
	}
}

func (s *Server) clientReader(cli *ClientHandler) (protocol.Transmission, error) {
	buffer, err := cli.transport.Read()
	if err != nil {
		return protocol.Transmission{}, err
	}

	transmission, err := protocol.UnmarshalTransmission(buffer)
	if err != nil {
		return protocol.Transmission{}, err
	}

	return transmission, nil
}

func (s *Server) giveCliNewId(client *ClientHandler) error {
	newId := protocol.GiveClientId{Id: client.id}
	_, trans, err := protocol.TransmissionFactory(protocol.GiveClientNewIdCode, newId)
	if err != nil {
		return err
	}

	return client.transport.Write(trans)
}

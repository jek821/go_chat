package main

import (
	"go_chat/Protocol"
	"go_chat/Shared"
	"go_chat/Utils"
	"net"
)

type Server struct {
	l           net.Listener
	cliHandlers map[int]ClientHandler
	Shared.ListenerLogic
	Shared.SenderLogic
	c chan Protocol.Payload
}

func newServer() *Server {
	conn, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		Utils.HandleErr(err)
	}
	newServer := Server{l: conn, cliHandlers: make(map[int]ClientHandler)}

	return &newServer
}

func (s *Server) acceptClients() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			Utils.HandleErr(err)
		}
		// TODO: Fix Client ID assignment
		NewCliHandler(conn, -1)

	}

}

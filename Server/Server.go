package main

import (
	"encoding/json"
	"fmt"
	"go_chat/Protocol"
	"go_chat/Shared"
	"go_chat/Utils"
	"net"
	"sync"
)

type Server struct {
	l            net.Listener
	HandlersLock sync.Mutex
	cliHandlers  map[int]*ClientHandler
	Shared.ListenerLogic
	Shared.SenderLogic
	c chan Protocol.Payload
}

func newServer() *Server {
	conn, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		Utils.HandleErr(err)
	}
	newServer := Server{l: conn, cliHandlers: make(map[int]*ClientHandler)}
	return &newServer
}

func (s *Server) acceptClients() {
	for {
		conn, err := s.l.Accept()
		if err != nil {
			Utils.HandleErr(err)
		}
		// TODO: Fix Client ID assignment
		newHandler := NewCliHandler(conn, -1, s.c)
		s.HandlersLock.Lock()
		s.cliHandlers[newHandler.CliId] = newHandler
		s.HandlersLock.Unlock()
	}
}

func (s *Server) PayloadParser(payload Protocol.Payload) {
	switch payload.Code {
	case Protocol.EndClientCode:
		var data Protocol.EndClient
		err := json.Unmarshal(payload.Data, &data)
		if err != nil {
			Utils.HandleErr(err)
		}
		s.EndClient(data.Id)
	default:
		fmt.Println("Unknown Payload Code in Server Payload Parser")
	}
}

func (s *Server) EndClient(CliId int) {
	s.HandlersLock.Lock()
	delete(s.cliHandlers, CliId)
	s.HandlersLock.Unlock()
}

func (s *Server) ChannelReader() {
	for payload := range s.c {
		s.PayloadParser(payload)
	}
}

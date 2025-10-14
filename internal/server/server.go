package server

import (
	"encoding/json"
	"fmt"
	"go_chat/pkg/protocol"
	"log/slog"
	"net"
)

type Server struct {
	Ip             string
	Port           int
	clients        map[int]*ClientHandler
	serverHandlers map[protocol.Code]protocol.ServerHandler
	clientPipe     chan protocol.Transmission
	sessions       map[int]*protocol.Session
	clientIdCount  int
}

func NewServer(ip string, port int) *Server {
	s := &Server{
		Ip:            ip,
		Port:          port,
		clients:       make(map[int]*ClientHandler),
		clientPipe:    make(chan protocol.Transmission, 100),
		sessions:      make(map[int]*protocol.Session),
		clientIdCount: 0,
	}
	s.registerHandlers()
	return s
}

func (s *Server) registerHandlers() {
	s.serverHandlers = make(map[protocol.Code]protocol.ServerHandler)

	s.serverHandlers[protocol.MsgCode] = &protocol.ServerMessageHandler{}
	s.serverHandlers[protocol.ConnectionRequestCode] = &protocol.ServerConnectionRequestHandler{}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	defer listener.Close()

	// Start message processor
	go s.processMessages()

	// Accept connections
	return s.acceptConnections(listener)
}

func (s *Server) acceptConnections(listener net.Listener) error {
	for {
		conn, err := listener.Accept()
		if err != nil {
			slog.Error("Error accepting connection", slog.Any("error", err))
			continue
		}

		newID := s.generateUID()
		newCliHandler := s.cliHandlerFactory(newID, conn)
		s.clients[newID] = &newCliHandler

		fmt.Printf("New client connected with ID: %d\n", newID)
		go s.cliHandler(&newCliHandler)
	}
}

func (s *Server) generateUID() int {
	s.clientIdCount++
	return s.clientIdCount
}

func (s *Server) processMessages() error {
	for trans := range s.clientPipe {
		handler, ok := s.serverHandlers[trans.Code]
		if !ok {
			slog.Warn("Unknown message code", slog.Int("code", int(trans.Code)))
			continue
		}

		if err := handler.Handle(trans, s); err != nil {
			slog.Error("Handler error",
				slog.Int("code", int(trans.Code)),
				slog.Any("error", err))
		}
	}
	return nil
}

// ServerContext interface implementations
func (s *Server) RouteToClient(clientID int, trans protocol.Transmission) error {
	client, ok := s.clients[clientID]
	if !ok {
		return fmt.Errorf("client %d not found", clientID)
	}

	data, err := json.Marshal(trans)
	if err != nil {
		return err
	}

	return client.transport.Write(data)
}

func (s *Server) BroadcastMessage(msg protocol.Message) error {
	_, data, err := protocol.TransmissionFactory(protocol.MsgCode, msg)
	if err != nil {
		return err
	}

	for id, client := range s.clients {
		if id == msg.OriginId {
			continue // Don't send back to sender
		}
		if err := client.transport.Write(data); err != nil {
			slog.Error("Broadcast failed",
				slog.Int("clientID", id),
				slog.Any("error", err))
		}
	}
	return nil
}

func (s *Server) GetClient(clientID int) bool {
	_, ok := s.clients[clientID]
	return ok
}

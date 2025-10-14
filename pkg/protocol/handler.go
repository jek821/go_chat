package protocol

import (
	"encoding/json"
	"log/slog"
)

type Handler interface {
	Handle(data []byte) error
}

type ServerHandler interface {
	Handle(trans Transmission, ctx ServerContext) error
}

type ServerContext interface {
	RouteToClient(clientId int, trans Transmission) error
	BroadCastMessage(msg Message) error
	GetClient(clientId int) bool
}

type ClientIDHandler struct {
	OnReceive func(id int)
}

func (h *ClientIDHandler) Handle(data []byte) error {
	var msg GiveClientId
	err := json.Unmarshal(data, &msg)
	if err != nil {
		slog.Error("unmarshal msg err:", err)
		return err
	}

	if h.OnReceive != nil {
		h.OnReceive(msg.Id)
	}
	return nil
}

type ConnectionRequestHandler struct {
	OnReceive func(req ConnectionRequest)
}

func (h *ConnectionRequestHandler) Handle(data []byte) error {
	var req ConnectionRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		slog.Error("unmarshal req err:", err)
		return err
	}

	if h.OnReceive != nil {
		h.OnReceive(req)
	}
	return nil
}

type MessageHandler struct {
	OnReceive func(msg Message)
}

func (h *MessageHandler) Handle(data []byte) error {
	var msg Message
	err := json.Unmarshal(data, &msg)
	if err != nil {
		slog.Error("unmarshal msg err:", err)
		return err
	}
	if h.OnReceive != nil {
		h.OnReceive(msg)
	}
	return nil
}

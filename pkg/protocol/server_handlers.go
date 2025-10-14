package protocol

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type ServerMessageHandler struct{}

func (h *ServerMessageHandler) Handle(trans Transmission, ctx ServerContext) error {
	var msg Message
	if err := json.Unmarshal(trans.Data, &msg); err != nil {
		slog.Error("Failed to unmarshal trans", err)
		return err
	}

	fmt.Printf("Message received: %v\n", msg)

	// server logic could broadcast
	return ctx.BroadCastMessage(msg)
}

// ServerConnectionRequestHandler Server handler for connection requests
type ServerConnectionRequestHandler struct{}

func (h *ServerConnectionRequestHandler) Handle(trans Transmission, ctx ServerContext) error {
	var req ConnectionRequest
	if err := json.Unmarshal(trans.Data, &req); err != nil {
		slog.Error("unmarshal req err:", err)
		return err
	}

	// Check if a target client exists
	if !ctx.GetClient(req.Target) {
		slog.Warn("Target client not found", slog.Int("target", req.Target))
		return fmt.Errorf("client %d not found", req.Target)
	}

	// Route the entire transmission to a target client
	return ctx.RouteToClient(req.Target, trans)
}

package Shared

import (
	"encoding/json"
	"fmt"
	"go_chat/Protocol"
	"go_chat/Utils"
	"log/slog"
	"net"
)

type ListenerLogic struct {
}

func (l *ListenerLogic) HandleIncomingPayLoads(conn net.Conn, handler func(payload Protocol.Payload)) {
	//TODO: FIX THIS BUFFER
	shittyBuffer := make([]byte, 1000)

	numBytes, err := conn.Read(shittyBuffer)
	if err != nil {
		Utils.HandleErr(err)
	}
	encodedTrans := shittyBuffer[:numBytes]
	var decodedTrans Protocol.Payload
	err = json.Unmarshal(encodedTrans, &decodedTrans)
	if err != nil {
		slog.Error("error reading incoming message", err)
		fmt.Println("ERROR UNMARSHALLING DATA")
	}
	handler(decodedTrans)
}

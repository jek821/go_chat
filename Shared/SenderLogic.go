package Shared

import (
	"go_chat/Protocol"
	"go_chat/Utils"
	"net"
)

type SenderLogic struct {
}

func (*SenderLogic) sendPayload(conn net.Conn, payload Protocol.Payload) {
	_, err := conn.Write(payload.EncodePayload())
	Utils.HandleErr(err)
}

func (*SenderLogic) SendAwait(conn net.Conn, payload Protocol.Payload) {
	_, err := conn.Write(payload.EncodePayload())
	Utils.HandleErr(err)
}

package Shared

import (
	"go_chat/Protocol"
	"go_chat/Utils"
	"net"
)

type SenderLogic struct {
}

func (*SenderLogic) SendPayload(conn net.Conn, payload Protocol.Payload) {
	_, err := conn.Write(payload.EncodePayload())
	if err != nil {
		Utils.HandleErr(err)
	}

}

func (*SenderLogic) SendAwait(conn net.Conn, payload Protocol.Payload) {
	_, err := conn.Write(payload.EncodePayload())
	if err != nil {
		Utils.HandleErr(err)
	}
}

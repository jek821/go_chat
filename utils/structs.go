package utils

import (
	"encoding/json"
	"net"
	"sync"
)

type Transmission struct {
	Uid  uint32          `json:"uid"`
	Code Code            `json:"code"`
	Data json.RawMessage `json:"data"`
}

type Message struct {
	OriginId  int    `json:"originid"`
	Body      string `json:"body"`
	SessionId int    `json:"session_id"`
}

type ConnectionRequest struct {
	Requester int `json:"requester"`
	Target    int `json:"target"`
}

type GiveClientId struct {
	Id int `json:"id"`
}

type Session struct {
	Id   int
	Cli1 int
	Cli2 int
}

type SessionAccept struct {
	AcceptorId int
	AcceptedId int
}

type SessionReject struct {
	RejectorId int
	RejectedId int
}

type Client struct {
	ClientID  int
	Conn      net.Conn
	Awaiters  map[uint32]chan Transmission
	AwaitLock sync.Mutex
}

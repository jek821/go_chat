package utils

import "encoding/json"

type Transmission struct {
	ID   int             `json:"id"`
	Code Code            `json:"code"`
	Data json.RawMessage `json:"data"`
}

type Message struct {
	Body      string `json:"body"`
	SessionId int    `json:"session_id"`
}

type ConnectionRequest struct {
	Requester int `json:"requester"`
	Target    int `json:"target"`
}

type GiveClientId struct {
	Code Code `json:"code"`
	Id   int  `json:"id"`
}

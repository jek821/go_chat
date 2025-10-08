package utils

import "encoding/json"

type Transmission struct {
	Code Code            `json:"code"`
	Data json.RawMessage `json:"data"`
}

type Registration struct {
}

type Message struct {
	Body      string `json:"body"`
	SessionId int    `json:"session_id"`
}

type GetSession struct {
	ClientId int `json:"client_id"`
}

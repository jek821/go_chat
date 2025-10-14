package utils

type Message struct {
	OriginId  int    `json:"originid"`
	Body      string `json:"body"`
	SessionId int    `json:"session_id"`
}

type ConnectionRequest struct {
	Requester int `json:"requester"`
	Target    int `json:"target"`
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

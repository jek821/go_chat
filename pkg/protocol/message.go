package protocol

// Message represents a chat message
type Message struct {
	OriginId  int    `json:"originid"`
	Body      string `json:"body"`
	SessionId int    `json:"session_id"`
}

// ConnectionRequest is sent when a client wants to connect to another
type ConnectionRequest struct {
	Requester int `json:"requester"`
	Target    int `json:"target"`
}

// GiveClientId is sent by server to assign a client an ID
type GiveClientId struct {
	Id int `json:"id"`
}

// Session represents a connection between two clients
type Session struct {
	Id   int
	Cli1 int
	Cli2 int
}

// SessionAccept is sent when a client accepts a connection
type SessionAccept struct {
	AcceptorId int `json:"acceptor_id"`
	AcceptedId int `json:"accepted_id"`
}

// SessionReject is sent when a client rejects a connection
type SessionReject struct {
	RejectorId int `json:"rejector_id"`
	RejectedId int `json:"rejected_id"`
}

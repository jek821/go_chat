package utils

type Transmission struct {
	Code Code
	Data interface{}
}

type Registration struct {
}

type Message struct {
	Body      string
	SessionId int
}

type getSession struct {
	ClientId int
}

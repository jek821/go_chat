package Protocol

import (
	"encoding/json"
	"go_chat/Utils"
)

type Payload struct {
	Code Code
	Data json.RawMessage
}

func (p *Payload) EncodePayload() json.RawMessage {
	encodedPayload, err := json.Marshal(p)
	Utils.HandleErr(err)
	return encodedPayload

}

type Test struct {
	FakeData string
}

type GiveClientId struct {
}

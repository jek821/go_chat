package protocol

import (
	"encoding/json"
	"fmt"
)

var uid uint32 = 99

type Transmission struct {
	Uid  uint32          `json:"uid"`
	Code Code            `json:"code"`
	Data json.RawMessage `json:"data"`
}

// UnmarshalTransmission decodes bytes into a Transmission
func UnmarshalTransmission(data []byte) (Transmission, error) {
	var trans Transmission
	err := json.Unmarshal(data, &trans)
	if err != nil {
		return Transmission{}, fmt.Errorf("error decoding transmission: %w", err)
	}
	return trans, nil
}

// NewTransmission creates a new transmission with encoded data
func NewTransmission(code Code, data any) (uint32, json.RawMessage, error) {
	uid++
	id := uid

	encodedData, err := json.Marshal(data)
	if err != nil {
		return 0, nil, err
	}

	trans := Transmission{
		Uid:  id,
		Code: code,
		Data: encodedData,
	}

	encodedTrans, err := json.Marshal(trans)
	if err != nil {
		return 0, nil, err
	}

	return id, encodedTrans, nil
}

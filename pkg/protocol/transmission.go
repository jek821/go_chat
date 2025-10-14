package protocol

import (
	"encoding/json"
	"fmt"
	"sync/atomic"
	"time"
)

var uid uint32 = 99

type Transmission struct {
	Uid  uint32          `json:"uid"`
	Code Code            `json:"code"`
	Data json.RawMessage `json:"data"`
}

func (t Transmission) Write(bytes []byte) error {
	//TODO implement me
	panic("implement me")
}

func (t Transmission) Close() error {
	//TODO implement me
	panic("implement me")
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

var uidCounter uint32

// generateUID creates a unique identifier for transmissions
func generateUID() uint32 {
	// Combine timestamp with atomic counter for uniqueness
	timestamp := uint32(time.Now().Unix())
	counter := atomic.AddUint32(&uidCounter, 1)
	return (timestamp << 16) | (counter & 0xFFFF)
}

// TransmissionFactory creates a new transmission with the given code and data
// Returns: (uid, marshaled transmission bytes, error)
func TransmissionFactory(code Code, data interface{}) (uint32, json.RawMessage, error) {
	// Generate unique ID for this transmission
	uid := generateUID()

	// Marshal the data payload
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to marshal data: %w", err)
	}

	// Create transmission structure
	trans := Transmission{
		Uid:  uid,
		Code: code,
		Data: dataBytes,
	}

	// Marshal the complete transmission
	transBytes, err := json.Marshal(trans)
	if err != nil {
		return 0, nil, fmt.Errorf("failed to marshal transmission: %w", err)
	}

	return uid, transBytes, nil
}

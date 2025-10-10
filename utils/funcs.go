package utils

import (
	"encoding/json"
	"fmt"
)

func UnmarshalData(msg []byte) (Transmission, error) {
	var decodedTransmission Transmission
	err := json.Unmarshal(msg, &decodedTransmission)
	if err != nil {
		return Transmission{}, fmt.Errorf("error decoding message: %w", err)
	}
	return decodedTransmission, nil
}

func MessageFactory(body string, sessionId int) Message {
	var newMessage = Message{Body: body, SessionId: sessionId}
	return newMessage
}

// Take in a defined struct and a code for that struct type
// Encode the struct and then put it inside of a new Transmission struct
// Then encode that entire transmission (also containing the code)
// Return the byte code for the entire transmission
func TransmissionFactory(code Code, data any, id int) (json.RawMessage, error) {
	encodedData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var newTransmission = Transmission{Code: code, Data: encodedData, ID: id}
	encodedTransmission, err := json.Marshal(newTransmission)
	if err != nil {
		return nil, err
	}
	return encodedTransmission, nil
}

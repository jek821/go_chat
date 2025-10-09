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

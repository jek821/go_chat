package client

import "C"
import (
	"encoding/json"
	"fmt"
	"go_chat/utils"
)

type Transmission struct {
	Uid  uint32          `json:"uid"`
	Code Code            `json:"code"`
	Data json.RawMessage `json:"data"`
}

func NewTransmission(uid uint32, code Code, data json.RawMessage) *Transmission {
	return &Transmission{
		Client: Client{},
	}
}

func (t *Transmission) TransmissionParser(transmission *Transmission) error {
	switch transmission.Code {
	case GiveClientNewIdCode:
		ClientLock.Lock()
		var data GiveClientId
		err := json.Unmarshal(transmission.Data, &data)
		if err != nil {
			return err
		}
		t.ClientID = data.Id
		ClientLock.Unlock()
	case ConnectionRequestCode:
	default:
		fmt.Printf("UNKNOWN CODE RECIEVED FROM SERVER")
	}
	return nil
}

func SendTransmission(transmission json.RawMessage) error {
	_, err := Client.Conn.Write(transmission)
	if err != nil {
		fmt.Println("Error Writing Message to Conn")
		return err
	}
	return nil
}

func SendAwait(id uint32, transmission json.RawMessage, channel chan Transmission) error {
	Client.AwaitLock.Lock()
	Client.Awaiters[id] = channel
	Client.AwaitLock.Unlock()
	err := sendTransmission(transmission)
	if err != nil {

		return err
	}
	return nil
}

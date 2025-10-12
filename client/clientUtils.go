package main

import (
	"bufio"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_chat/utils"
	"io"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
)

var ClientLock sync.Mutex
var Client = utils.Client{}

func StartClient() error {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("Error Connecting to server in StartClient()")
		return err
	}
	defer conn.Close()
	ClientLock.Lock()
	Client.Conn = conn
	Client.Awaiters = make(map[uint32]chan utils.Transmission)
	ClientLock.Unlock()
	return nil
}

func Listener() error {
	var length uint32
	for {
		binary.Read(Client.Conn, binary.BigEndian, &length)
		buffer := make([]byte, length)
		_, err := io.ReadFull(Client.Conn, buffer)
		if err != nil {
			return err
		}
		encodedMsg := buffer[0:length]
		transmission, err := utils.UnmarshalData(encodedMsg)
		if err != nil {
			return err
		}
		Client.AwaitLock.Lock()
		_, ok := Client.Awaiters[transmission.Uid]
		if ok {
			Client.Awaiters[transmission.Uid] <- transmission
			delete(Client.Awaiters, transmission.Uid)
		}
		TransmissionParser(transmission)
		Client.AwaitLock.Unlock()
	}
}

func TransmissionParser(transmission utils.Transmission) error {
	switch transmission.Code {
	case utils.GiveClientNewIdCode:
		ClientLock.Lock()
		var data utils.GiveClientId
		err := json.Unmarshal(transmission.Data, &data)
		if err != nil {
			return err
		}
		Client.ClientID = data.Id
		ClientLock.Unlock()
	case utils.ConnectionRequestCode:

	default:
		fmt.Printf("UNKNOWN CODE RECIEVED FROM SERVER")
	}
	return nil
}

func sendTransmission(transmission json.RawMessage) error {
	_, err := Client.Conn.Write(transmission)
	if err != nil {
		fmt.Println("Error Writing Message to Conn")
		return err
	}
	return nil
}

func initializeMenu() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Menu Options:\nType (X) -> Initiate Connection")
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		text = strings.TrimSuffix(text, "\n")
		if strings.EqualFold("x", text) {
			fmt.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
			fmt.Println("Enter Client ID You Would Like To Send Request To: ")
			text, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			text = strings.TrimSuffix(text, "\n")
			targetId, err := strconv.Atoi(text)
			if err != nil {
				return err
			}
			GetReqResp := make(chan utils.Transmission, 1)
			request := utils.ConnectionRequest{Requester: Client.ClientID, Target: targetId}
			transId, trans, err := utils.TransmissionFactory(utils.ConnectionRequestCode, request)
			if err != nil {
				return err
			}
			SendAwait(transId, trans, GetReqResp)
			response := <-GetReqResp
			if response.Code == utils.SessionRequestAcceptedCode {
				fmt.Println("Request Rejected")
			}
			if response.Code == utils.SessionRequestRejectedCode {
				fmt.Println("Request Rejected")
			}
		}
	}
}

func GetReqResp(req utils.ConnectionRequest) error {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Incoming Connection Request From Client: %d\nType (y) to accept to anything else to reject!", req.Requester)
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		text = strings.TrimSuffix(text, "\n")
		if strings.EqualFold("y", text) {
		}
	}

}

func SendAwait(id uint32, transmission json.RawMessage, channel chan utils.Transmission) error {
	Client.AwaitLock.Lock()
	Client.Awaiters[id] = channel
	Client.AwaitLock.Unlock()
	err := sendTransmission(transmission)
	if err != nil {
		return err
	}
	return nil
}

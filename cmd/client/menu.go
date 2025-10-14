package main

import (
	"bufio"
	"fmt"
	"go_chat/utils"
	"os"
	"strconv"
	"strings"
)

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

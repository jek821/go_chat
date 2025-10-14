package client

import (
	"bufio"
	"fmt"
	"go_chat/utils"
	"os"
	"strings"
)

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

import (
	"errors"
	"fmt"
	"net"
	"time"
)

var ClientId int

func main() {
	// conn, err := startCli()
	_, err := startCli()
	if err != nil {
		return
	}
	// RegisterClient(conn)
	for {
		time.Sleep((1000000000000000))
	}

}

func startCli() (net.Conn, error) {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		return nil, errors.New("ERROR CONNECTING TO SERVER")
	}

	defer conn.Close()

	fmt.Println("Connection To Server Successful!")

	return conn, nil
}

// func sendMsg(conn net.Conn, sessionId int) {
// 	for {
// 		fmt.Println("input text:")
// 		reader := bufio.NewReader(os.Stdin)
// 		msg, err := reader.ReadString('\n')
// 		if err != nil {
// 			fmt.Println("Error Reading Input")
// 		}
// 		var msg = utils.Message{}
// 		data := encodeData(msg)
// 		_, err := conn.Write(data)
// 		if err != nil {
// 			fmt.Println("Error Writing Message to Conn")
// 			return
// 		}
// 	}
// }

// func encodeData(data utils.Transmission) []byte {
// 	EncodedTransaction, err := json.Marshal(data)
// 	if err != nil {
// 		fmt.Printf("Error Encoding Message: %s", data)
// 	}
// 	return EncodedTransaction

// }

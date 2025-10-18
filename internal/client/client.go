package client

import (
	"fmt"
	"go_chat/pkg/protocol" // ← Changed from utils
	"net"
	"sync"
)

type Client struct {
	transport Transport
	clientID  int
	conn      net.Conn
	awaiters  map[uint32]chan protocol.Transmission
	awaitLock sync.Mutex
	handlers  map[protocol.Code]protocol.Handler
}

func NewClient(serverAddr string) (*Client, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	client := &Client{
		conn:      conn,
		transport: &TCPTransport{Conn: conn},
		awaiters:  make(map[uint32]chan protocol.Transmission),
		handlers:  make(map[protocol.Code]protocol.Handler),
	}

	client.registerHandler()

	return client, nil
}

func (c *Client) registerHandler() {
	// handle client ID assignment
	c.handlers[protocol.GiveClientNewIdCode] = &protocol.ClientIDHandler{
		OnReceive: func(id int) {
			c.SetID(id)
			fmt.Printf("Received new client id %d\n", id)
		},
	}

	c.handlers[protocol.ConnectionRequestCode] = &protocol.ConnectionRequestHandler{
		OnReceive: func(req protocol.ConnectionRequest) {
			fmt.Printf("Received connection request from client %d\n", req.Requester)
		},
	}

	c.handlers[protocol.MsgCode] = &protocol.MessageHandler{
		OnReceive: func(msg protocol.Message) {
			fmt.Printf("[Client %d received message] %s\n", msg.OriginId, msg.Body)
		},
	}
}

// GetID returns the client's ID
func (c *Client) GetID() int {
	return c.clientID
}

// SetID sets the client's ID
func (c *Client) SetID(id int) {
	c.clientID = id
}

// Close closes the client connection
func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// GetConn returns the connection
func (c *Client) GetConn() net.Conn {
	return c.conn
}

// RegisterAwaiter registers a channel waiting for a specific transmission
func (c *Client) RegisterAwaiter(uid uint32, ch chan protocol.Transmission) {
	c.awaitLock.Lock()
	defer c.awaitLock.Unlock()
	c.awaiters[uid] = ch
}

// GetAwaiter retrieves and removes an awaiter channel
func (c *Client) GetAwaiter(uid uint32) (chan protocol.Transmission, bool) {
	c.awaitLock.Lock()
	defer c.awaitLock.Unlock()
	ch, ok := c.awaiters[uid]
	if ok {
		delete(c.awaiters, uid)
	}
	return ch, ok
}

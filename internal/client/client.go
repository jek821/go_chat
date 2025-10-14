package client

import (
	"fmt"
	"go_chat/pkg/protocol" // ← Changed from utils
	"net"
	"sync"
)

type Client struct {
	clientID  int
	conn      net.Conn
	awaiters  map[uint32]chan protocol.Transmission // ← Changed
	awaitLock sync.Mutex
}

func NewClient(serverAddr string) (*Client, error) {
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to server: %w", err)
	}

	client := &Client{
		conn:     conn,
		awaiters: make(map[uint32]chan protocol.Transmission), // ← Changed
	}

	return client, nil
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

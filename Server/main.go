package Server

import "net"

type Server struct {
	conn        net.Conn
	cliHandlers map[int]ClientHandler
}

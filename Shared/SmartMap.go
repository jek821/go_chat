package Shared

import (
	"go_chat/Protocol"
	"sync"
)

type AwaitMap struct {
	Map    map[int]Protocol.Payload
	chLock sync.RWMutex
	ch     chan string
}

func (m *AwaitMap) NewWatcher()

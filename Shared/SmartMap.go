package Shared

import (
	"go_chat/Protocol"
	"sync"
)

type AwaitMap struct {
	Map    map[int]chan Protocol.Payload
	chLock sync.Mutex
}

func CreateAwaitMap() AwaitMap {
	return AwaitMap{Map: make(map[int]chan Protocol.Payload),
		chLock: sync.Mutex{}}
}

func (m *AwaitMap) NewAwaiter(pid int, ch chan Protocol.Payload) {
	m.chLock.Lock()
	m.Map[pid] = ch
	m.chLock.Unlock()
}

func (m *AwaitMap) ResolveWaiter(payload Protocol.Payload) {
	m.chLock.Lock()
	ch := m.Map[payload.Pid]
	ch <- payload
	delete(m.Map, payload.Pid)
	m.chLock.Unlock()
}

package Shared

import (
	"go_chat/Protocol"
	"sync"
)

type AwaitMap struct {
	Map    map[int]chan Protocol.Payload
	chLock sync.Mutex
}

func (m *AwaitMap) NewAwaiter(pid int, ch chan Protocol.Payload) {
	m.chLock.Lock()
	m.Map[pid] = ch
	m.chLock.Unlock()
}

func (m *AwaitMap) ResolveWaiter(pid int, payload Protocol.Payload) {
	m.chLock.Lock()
	ch := m.Map[pid]
	ch <- payload
	delete(m.Map, pid)
	m.chLock.Unlock()
}

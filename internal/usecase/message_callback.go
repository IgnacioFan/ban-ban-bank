package usecase

import "sync"

type MessageCallback struct {
	Lock  sync.RWMutex
	Chans map[string]chan []byte
}

func (cb *MessageCallback) Set(key string, data chan []byte) {
	cb.Lock.Lock()
	defer cb.Lock.Unlock()
	cb.Chans[key] = data
}

func (cb *MessageCallback) Delete(key string) {
	cb.Lock.Lock()
	defer cb.Lock.Unlock()
	delete(cb.Chans, key)
}

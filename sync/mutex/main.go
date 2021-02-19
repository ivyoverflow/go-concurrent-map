package main

import (
	"fmt"
	"sync"
)

type concurrentMap struct {
	mu *sync.RWMutex
	wg *sync.WaitGroup
}

func (m *concurrentMap) get(h map[string]string, key string) {
	defer m.wg.Done()
	m.mu.RLock()
	value, ok := h[key]
	if !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println(value)
	}

	m.mu.RUnlock()
}

func (m *concurrentMap) add(h map[string]string, key string, value string) {
	defer m.wg.Done()
	m.mu.Lock()
	h[key] = value
	m.mu.Unlock()
}

func (m *concurrentMap) delete(h map[string]string, key string) {
	defer m.wg.Done()
	m.mu.Lock()
	delete(h, key)
	m.mu.Unlock()
}

func main() {
	mu := &sync.RWMutex{}
	wg := &sync.WaitGroup{}
	m := concurrentMap{mu, wg}
	h := make(map[string]string)
	clients := [8]string{"Bob", "Jack", "John", "Bill", "Alice", "Adam", "Garry", "Howard"}
	for _, client := range clients {
		wg.Add(5)
		go m.get(h, client)
		go m.add(h, client, "Hello World!")
		go m.get(h, client)
		go m.delete(h, client)
		go m.get(h, client)
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
)

type concurrentMap struct {
	wg     *sync.WaitGroup
	locked chan struct{}
}

func (m *concurrentMap) get(h map[string]string, key string) {
	defer m.wg.Done()
	m.locked <- struct{}{}
	value, ok := h[key]
	if !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println(value)
	}

	<-m.locked
}

func (m *concurrentMap) add(h map[string]string, key string, value string) {
	defer m.wg.Done()
	m.locked <- struct{}{}
	h[key] = value
	<-m.locked
}

func (m *concurrentMap) delete(h map[string]string, key string) {
	defer m.wg.Done()
	m.locked <- struct{}{}
	delete(h, key)
	<-m.locked
}

func main() {
	ch := make(chan struct{}, 1)
	wg := &sync.WaitGroup{}
	m := concurrentMap{wg, ch}
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

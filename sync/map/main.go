package main

import (
	"fmt"
	"sync"
)

type concurrentMap struct {
	m  *sync.Map
	wg *sync.WaitGroup
}

func (m *concurrentMap) get(key string) {
	defer m.wg.Done()
	value, ok := m.m.Load(key)
	if !ok {
		fmt.Println("Key not found")
	} else {
		fmt.Println(value)
	}
}

func (m *concurrentMap) add(key string, value string) {
	defer m.wg.Done()
	m.m.Store(key, value)
}

func (m *concurrentMap) delete(key string) {
	defer m.wg.Done()
	m.m.Delete(key)
}

func main() {
	m := &sync.Map{}
	wg := &sync.WaitGroup{}
	cm := &concurrentMap{m, wg}
	clients := [8]string{"Bob", "Jack", "John", "Bill", "Alice", "Adam", "Garry", "Howard"}
	for _, client := range clients {
		wg.Add(5)
		go cm.get(client)
		go cm.add(client, "Hello World!")
		go cm.get(client)
		go cm.delete(client)
		go cm.get(client)
	}

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var (
		lock sync.RWMutex
		m    = make(map[int]int)
	)
	lock.RLock()

	defer func() {
		lock.RUnlock()
		fmt.Println("RUnlock Unlock")
	}()
	for i := 0; i < 10; i++ {
		j := i
		go func(i int) {
			fmt.Println("Rlock")
			lock.RLock()
			fmt.Println("Rlock ok")
			m[i] = i

			fmt.Println("RUnlock Unlock")
		}(j)
	}

	for i := 0; i < 10; i++ {
		j := i
		go func(i int) {
			fmt.Println("wlock ping")
			lock.Lock()
			fmt.Println("wlock ok")
			m[i] = i
			lock.Unlock()
			fmt.Println("wlock Unlock")
		}(j)
	}
	time.Sleep(time.Second * 5)
	for k, v := range m {
		fmt.Println(k, v)
	}
}

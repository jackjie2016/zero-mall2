package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex

	chrlockTwice := make(chan struct{}) // Used to control rlockTwice
	rlockTwice := func() {
		mu.RLock()
		fmt.Println("first Rlock succeeded")
		<-chrlockTwice
		<-chrlockTwice
		fmt.Println("trying to Rlock again")
		mu.RLock()
		fmt.Println("second Rlock succeeded")
		mu.RUnlock()
		mu.RUnlock()
	}

	chLock := make(chan struct{}) // Used to contol lock
	lock := func() {
		<-chLock
		fmt.Println("about to Lock")
		mu.Lock()
		fmt.Println("Lock succeeded")
		mu.Unlock()
		<-chLock
	}

	control := func() {
		chrlockTwice <- struct{}{}
		chLock <- struct{}{}

		close(chrlockTwice)
		close(chLock)
	}

	go control()
	go lock()
	rlockTwice()
}

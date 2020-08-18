package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ticket int = 10
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(4)

	go saleTickets("window 1")
	go saleTickets("window 2")
	go saleTickets("window 3")
	go saleTickets("window 4")
	// time.Sleep(5 * time.Second)
	wg.Wait()
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	defer wg.Done()
	for {
		mutex.Lock()
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "sale", ticket)
			ticket--
		} else {
			mutex.Unlock()
			fmt.Println(name, "no ticket")
			break
		}
		mutex.Unlock()
	}
}

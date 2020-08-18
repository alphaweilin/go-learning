package main

import (
	"fmt"
	"time"
)

func main() {
	a := 1
	go func() {
		a = 2
		fmt.Println("value of a on goroutine", a)
	}()

	a = 3
	time.Sleep(1)
	fmt.Println("value of a on main goroutine", a)
}

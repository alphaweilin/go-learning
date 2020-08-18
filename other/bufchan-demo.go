package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		The buffer size is the number of elements that can be sent to the channel without the send blocking.
		By default, a channel has a buffer size of 0 (you get this with make(chan int)).
		This means that every single send will block until another goroutine receives from the channel.
	*/
	ch := make(chan string, 4)
	go sendData(ch)

	for {
		v, ok := <-ch
		if ok {
			fmt.Println("\t读取的数据是:", v, ok)
		} else {
			fmt.Println("读完了。。。", ok)
			break
		}
	}
	fmt.Println("main....over....")
}

func sendData(ch chan string) {
	for i := 1; i <= 10; i++ {
		ch <- "data:" + strconv.Itoa(i)
		fmt.Printf("子goroutine中写出第 %d 个数据\n", i)
	}
	close(ch)
}

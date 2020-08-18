package main

import "fmt"

func main() {
	funcB()
}

func funcB() {
	defer func() {
		if msg := recover(); msg != nil { //相当于catch，捕获异常进行处理，让程序继续运行
			fmt.Println("get painc:", msg)
			fmt.Println("recover program")
		}
	}()
	panic("panic test")
}

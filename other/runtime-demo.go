package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GOROOT--->", runtime.GOROOT())
	fmt.Println("os/platform--->", runtime.GOOS)
	fmt.Println("cpu number--->", runtime.NumCPU())
}

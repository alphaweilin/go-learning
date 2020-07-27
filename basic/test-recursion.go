package main

import "fmt"

func sum(x int) int {
	if x == 1 {
		return x
	}

	return x + sum(x-1)
}

func main() {
	fmt.Println(sum(10))
}

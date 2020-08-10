package main

import "fmt"

type data [2]int

func main() {
	switch x := 5; x {
	default:
		fmt.Println(x)
	case 5:
		x += 10
		fmt.Println(x)
		//穿透case
		fallthrough
	case 6:
		x += 20
		fmt.Println(x)
	}
}

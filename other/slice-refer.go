package main

import "fmt"

func main() {
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	pass(slice)
	fmt.Println(slice)
}

func pass(slice []int) {
	fmt.Println("slice on pass", slice)
	slice = []int{2, 2, 2, 2, 2, 2}
	fmt.Println("slice on pass", slice)
}

package main

import "fmt"

func main() {
	a := 10
LOOP:
	for a < 20 {
		if a == 15 {
			a++
			goto LOOP
		}

		fmt.Printf("value of a is: %d\n", a)
		a++
	}
}

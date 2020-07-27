package main

import (
	"fmt"
)

type MyInt int

func (f MyInt) Abs() int {
	if f < 0 {
		return int(-f)
	}
	return int(f)
}

func main() {
	i := MyInt(-5)
	fmt.Println(i.Abs())
}

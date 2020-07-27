package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

//方法就是一类带特殊的 接收者 参数的函数
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

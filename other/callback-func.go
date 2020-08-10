package main

import "fmt"

func main() {
	/*
		高阶函数：
			根据go语言的数据类型特点，可以将一个函数作为另一个函数的参数
		fun1(),fun2()
		将fun1作为fun2的参数。
			fun2: 就叫高阶函数
				接收了一个函数作为参数的函数，高阶函数
			fun1: 回调函数
				作为另外一个函数参数的函数，叫做回调函数
	*/

	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", oper)

	res1 := add(1, 2)
	fmt.Println(res1)

	res2 := oper(10, 20, add)
	fmt.Println(res2)

	res3 := oper(5, 2, sub)
	fmt.Println(res3)

	//匿名函数
	fun1 := func(a, b int) int {
		return a * b
	}
	res4 := oper(10, 4, fun1)
	fmt.Println(res4)

	res5 := oper(100, 8, func(a, b int) int {
		if b == 0 {
			fmt.Println("除数不能为0")
			return 0
		}
		return a / b
	})
	fmt.Println(res5)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}

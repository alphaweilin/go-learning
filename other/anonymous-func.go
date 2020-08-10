package main

import "fmt"

func main() {
	/*
		匿名函数通常只使用一次
		Go语言支持函数式编程
		1.将匿名函数作为另外一个函数的参数，回调函数
		2.将匿名函数作为另外一个函数的返回值，闭包
	*/
	func() {
		fmt.Println("我是一个匿名函数")
	}()

	fun3 := func() {
		fmt.Println("我也是一个匿名函数")
	}
	fun3()

	//定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	//定义带返回值的匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(10, 20) //匿名函数调用了，将执行结果返回res1
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a + b
	} //将匿名函数的值，赋值给res2
	fmt.Println(res2)

	fmt.Println(res2(30, 40))
}

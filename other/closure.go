package main

import "fmt"

func main() {
	/*
		go语言支持函数式编程
			支持一个函数作为另外一个函数的参数
			也支持将一个函数作为另外一个函数的返回值

		闭包(closure):
			一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量(外层函数中的参数，或外层函数中直接定义的变量)，并且该外层函数的返回值就是这个内层函数。

			这个内层函数和外层函数的局部变量，统称为闭包结构。

			局部变量的生命周期会发生改变，正常的局部变量随着函数调用而创建，随着函数结束而销毁;
			但是闭包结构中的外层函数的局部变量并不会随着外层函数调用结束而销毁，因为内层函数还要继续使用。
	*/

	res1 := increment()      //res1 = fun
	fmt.Printf("%T\n", res1) //func() int
	fmt.Println(res1)
	v1 := res1()
	fmt.Println(v1)
	v2 := res1()
	fmt.Println(v2)

	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())

	res2 := increment() //每次调用increment，都会产生新的局部变量i
	fmt.Println(res2)
	v3 := res2()
	fmt.Println(v3) //1
}

func increment() func() int { //外层函数
	//1.定义一个局部变量
	i := 0
	//2.定义一个匿名函数，给变量自增并返回
	fun := func() int { //内层函数
		i++
		return i
	}
	//3.返回匿名函数
	return fun
}

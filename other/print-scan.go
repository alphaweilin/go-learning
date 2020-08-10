package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
输入和输出：
	fmt包：输入，输出

		输出：
			Print() //打印
			Printf() //格式化打印
			Println() //换行打印

		格式化打印占位符：
		%v 原样打印
		%T 打印类型
		%t bool类型
		%s 字符串
		%f 浮点
		%d 10进制整数
		%o 8进制
		%x,%X 16进制
			%x: 0-9, a-f
			%X: 0-9, A-F
		%c 打印字符
		%p 打印地址
		......

		输入：
		Scanln()
		Scanf()

	bufio包
*/
func main() {
	fmt.Println("hello world")

	i := 32
	fmt.Printf("type is %T, address is %p, value is %v\n", i, &i, i)
	fmt.Println(&i)
	fmt.Println("-------------------------------------------------------------")

	/*var x int
	var y float64
	fmt.Println("please input two number:")
	fmt.Scanln(&x, &y) //读取键盘输入，通过操作地址，赋值给x和y  阻塞式
	fmt.Println(x, y)

	fmt.Scanf("%d,%f", &x, &y)
	fmt.Printf("%d,%f", x, y)*/

	fmt.Println("please input a string")
	reader := bufio.NewReader(os.Stdin)
	s1, err := reader.ReadString('\n') //读入以换行符结尾的输入
	if err == nil {
		fmt.Println("Read string: ", s1)
	}

}

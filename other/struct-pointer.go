package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p1 := Person{name: "zhangsan", age: 18}
	fmt.Println(p1)
	fmt.Printf("%p,%T\n", &p1, p1)

	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p,%T\n", pp1, pp1)
	fmt.Printf("%p,%T\n", &pp1, pp1)

	//使用内置函数new(),go语言中专门用于创建某种类型的指针的函数
	pp2 := new(Person)
	fmt.Println(pp2)
	fmt.Printf("%T\n", pp2)
	pp2.name = "lisi"
	pp2.age = 25
	fmt.Println(*pp2)

	pp3 := new(int)
	fmt.Println(pp3)
	fmt.Println(*pp3)
}

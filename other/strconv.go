package main

import (
	"fmt"
	"strconv"
)

func main() {
	//1.bool类型
	s1 := "true"
	b1, err := strconv.ParseBool(s1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T, %t\n", b1, b1)

	ss1 := strconv.FormatBool(b1)
	fmt.Printf("%T, %s\n", ss1, ss1)

	//2.整数
	s2 := "100"
	i2, err := strconv.ParseInt(s2, 10, 64) //10进制，最大长度为64
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T, %d\n", i2, i2)

	ss2 := strconv.FormatInt(i2, 10) //10进制
	fmt.Printf("%T, %s\n", ss2, ss2)

	//itoa(),atoi()
	i3, _ := strconv.Atoi("-42")
	fmt.Printf("%T, %d\n", i3, i3)
	ss3 := strconv.Itoa(-42)
	fmt.Printf("%T, %s\n", ss3, ss3)
}

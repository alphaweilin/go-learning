package main

import "fmt"

func main() {
	//字符串是字节的切片
	slice1 := []byte{65, 66, 67, 68, 69}
	s1 := string(slice1)
	// 字符串不允许修改，执行下行代码会报错：cannot assign to s1[0]
	// s1[0] = 70
	fmt.Println(s1)

	s2 := "abcde"
	slice2 := []byte(s2)
	fmt.Println(slice2)
}

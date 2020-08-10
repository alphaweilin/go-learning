package main

import "fmt"

func main() {
	var arr = [5]int{1, 2, 3, 4, 5}
	//数组是值类型
	fmt.Printf("address of array is: %p\n", &arr)
	fmt.Printf("address of array is: %p\n", &arr[0])

	//切片是引用类型
	//切片底层是数组，添加元素超过容量时候，切片会扩容
	var slice = []int{1, 2, 3, 4, 5}
	fmt.Printf("address of slice is: %p, length is: %d, capacity is %d\n", slice, len(slice), cap(slice))
	slice = append(slice, 7, 8, 9)
	//扩容后，切片地址会改变
	fmt.Printf("address of slice is: %p, length is: %d, capacity is %d\n", slice, len(slice), cap(slice))

	s1 := []int{1, 2, 3, 4, 5}
	s2 := s1
	fmt.Printf("%p\n", s1)
	fmt.Printf("%p\n", s2)
	fmt.Printf("%p\n", &s1)
	fmt.Printf("%p\n", &s2)
}

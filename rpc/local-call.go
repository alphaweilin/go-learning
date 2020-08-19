package main

import (
	"fmt"
	"math"
)

//数学计算
type MathUtil struct {
}

//该方法向外暴露：提供计算圆形面积的服务
func (mu *MathUtil) CalculateCircleArea(req float32, resp *float32) error {
	*resp = math.Pi * req * req //圆形的面积 s = π * r * r
	return nil                  //返回类型
}

func main() {
	var req float32 = 2
	var resp *float32
	var defaultVaule float32 = 0
	resp = &defaultVaule
	mathUtil := MathUtil{}
	mathUtil.CalculateCircleArea(req, resp)
	fmt.Println(*resp)
}

package main

import (
	"fmt"
)

// 浮点数
func main() {
	f1 := 1.23456
	fmt.Printf("%T\n", f1)  // 使用%T判断浮点数属于什么类型 输出结果为float64 go语言中float默认都是float64位
	f2 := float32(1.234567) //使用float32定义32位浮点数
	fmt.Printf("%T\n", f2)
}

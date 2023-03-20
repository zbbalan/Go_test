package main

import (
	"fmt"
)

// fmt 占位符
func main() {
	// 查看类型
	a1 := 1.23456456
	fmt.Printf("%T\n", a1) //查看类型
	fmt.Printf("%v\n", a1) //查看变量值
	fmt.Printf("%d\n", a1) //查看二进制
	fmt.Printf("%b\n", a1) //查看十进制
	fmt.Printf("%o\n", a1) //查看八进制
	fmt.Printf("%x\n", a1) //查看十六进制
	a2 := "hello"
	fmt.Printf("%T,%s\n", a2, a2) // 查看字符串
}

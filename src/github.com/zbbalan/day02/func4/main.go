package main

import (
	"fmt"
)

// 匿名函数

func main() {
	// 匿名函数用于在函数内部无法声明带名字的函数
	f1 := func(x, y int) {
		fmt.Println(x + y)
	}
	f1(10, 20)
	// 如果只是一次调用函数可以写成立即执行函数
	func(x, y int) {
		fmt.Println("hello,world")
		fmt.Println(x + y)
	}(100, 200) // 函数后面带括号立即执行
}

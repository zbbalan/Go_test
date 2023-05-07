package main

import (
	"fmt"
)

// 变量的作用域
var x = 1000 // 定义一个全局变量

func f1() {
	// 函数中查找变量的顺序
	// 先在函数内部查找
	// 找不到就往外面查找
	fmt.Println(x)
}
func main() {
	f1()
}

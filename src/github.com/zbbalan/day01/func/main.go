package main

import (
	"fmt"
)

// 函数是一段代码的封装，把一个逻辑抽象出来分装到一个函数中给他起个名字，每次用到直接用函数名调用
// 使用函数让代码更清晰，更简洁
func sum(x int, y int) (ret int) { // ret 是函数的名称
	return x + y
}

// 多个返回值
func Alan() (int, string) {
	return 666, "中国"
}

// 可变长参数，必须放在函数参数的最后
func Zbb(a string, b ...int) {
	fmt.Printf(a)
	fmt.Println(b)

}
func main() {
	// 函数
	r := sum(1, 2)
	fmt.Println(r)
	m, n := Alan()
	fmt.Println(m, n)
	Zbb("沙河")
	Zbb("中国", 1, 2, 3, 4, 4, 4, 4, 24, 21, 4, 1)
}

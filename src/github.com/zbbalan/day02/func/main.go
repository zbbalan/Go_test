package main

import (
	"fmt"
)

// 函数：一段代码的封装
func f1() {
	fmt.Println("hello,Zbb")

}

// 函数传参
func f2(name string) {
	fmt.Println("HELLO", name)
}

// 可变参数
func f3(x, y int) int {
	sum := x + y
	return sum
}

// 可变参数2
func f4(j string, k ...int) int {
	fmt.Println(j, k) // j k 是一个 int 类型的切片
	return 1
}

// 命名返回值
func f5(x, y int) (sum int) {
	sum = x + y // 如果使用命名返回值，那么在函数中可以直接使用返回值变量
	return      // return 后面可以省略返回值变量
}

// Go 语言支持多个返回值
func f6(x, y int) (sum int, sub int) {
	sum = x + y
	sub = x - y
	return
}
func main() {
	f1()      // 调用上面封装的F1 函数
	f2("ZBB") // 调用上面的传参
	fmt.Println(f3(100, 200))
	f4("沙河", 1, 2, 3, 4, 5)
	fmt.Println(f5(20, 30))
	fmt.Println(f6(60, 40))
}

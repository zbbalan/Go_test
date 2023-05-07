package main

import (
	"fmt"
)

// 函数类型
func f1() {
	fmt.Println("hello,中国")
}
func f2() int {
	return 1000
}

// 函数也可以作为参数类型  只要满足参数类型都可调用
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}

func ff(a, b int) int {
	return a + b
}

// 函数也可以作为返回值
func f4(x func() int) func(int, int) int {
	return ff
}
func main() {
	a := f1
	fmt.Printf("%T\n", a) // 返回为函数类型 ，前面没有定义 返回为 func ()
	b := f2
	fmt.Printf("%T\n", b) // 返回为func（）int
	f3(f2)
	f7 := f4(f2)
	fmt.Printf("%T\n", f7)
}

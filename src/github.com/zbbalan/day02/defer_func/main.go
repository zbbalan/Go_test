package main

import (
	"fmt"
)

// defer函数多用于资源释放
func deferDemon() {
	fmt.Println("start")
	defer fmt.Println("延迟处理")    // defer 把他后面的语句延迟到函数即将返回的时候在执行
	defer fmt.Println("延迟处理222") // 一个函数中可以有多个defer 语句
	defer fmt.Println("延迟处理333") // defer 执行顺序按照先进后出顺序延迟
	fmt.Println("end")
}
func main() {
	deferDemon()
}

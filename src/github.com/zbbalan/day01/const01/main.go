package main

import (
	"fmt"
)

// 常量，定义后不能修改，程序运行期间不会改变
const pi = 3.1415926535

// 批量声明常量
const (
	statusok = 200
	notfound = 400
)

// 批量声明常量后没有写值则以第一个常量值为准
const (
	n1 = 100
	n2
	n3
	n4
)

// iota 初始化为0 在常量中没新增一次则加1
const (
	a1 = iota // 0
	a2
	a3
)

func main() {
	fmt.Print(n1)
	fmt.Print(n2)
	fmt.Print(n3)
	fmt.Print(n4)
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
}

// day10常量和iota

package main

import (
	"fmt"
)

// 声明变量(全局变量)声明可以不使用
var (
	name string
	age  int
	isok bool
)

// 作用域
func main() {
	//变量赋值必须使用
	name = "zbb"
	age = 26
	isok = true
	//输出变量
	fmt.Printf("name:%s\n", name)
	fmt.Print(isok)
	fmt.Println(age)
	// 声明变量并赋值
	var s1 string = "Alan"
	fmt.Printf(s1)
	// 类型推导（根据值判断改变量是什么类型）
	var s2 = 100
	fmt.Println(s2)
	// 简短变量声明
	s3 := "hello"
	fmt.Printf(s3)
}

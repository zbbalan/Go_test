package main

import (
	"fmt"
)

func main() {
	// &: 取内存地址
	// *: 根据内存地址取值
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
	m := *p
	fmt.Println(m)
	fmt.Printf("%T\n", m)
	// new 函数申请一个内存地址
	// make 和new都是用来申请内存的
	//new 一般用的少 ，给基础数据类型申请内存 string，int
	// make 用来给slice map chan 申请内存 make函数返回的是这三个类型对应的本身
	var a = new(int)
	*a = 1000
	fmt.Println(*a)

}

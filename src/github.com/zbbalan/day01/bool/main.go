package main

import (
	"fmt"
)

// 布尔值默认是false
// 布尔值数据类型只有true 和false
func main() {
	b1 := true
	var b2 bool
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T value:%v\n", b2, b2) //value:%v 打印变量的值 输出结果为false

}

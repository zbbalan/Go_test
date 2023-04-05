// for 循环
package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i)
	}
	// 也可以先定义在循环
	var n = 5
	for ; n < 15; n++ {
		fmt.Println(n)
	}
	// for range 循环打印索引
	s := "helloword"
	for i, v := range s { //i代表索引，将s字符赋值给v
		fmt.Printf("%d %c\n", i, v)
	}
}

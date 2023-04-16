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
	// go 语言中for流程控制
	for i := 0; i < 10; i++ {
		if i == 10 {
			break //当i =10 时跳出循环
		}
		fmt.Println(i)
	}
	fmt.Printf("break循环跳过")
	// 跳过打印语句
	for i := 4; i < 8; i++ {
		if i == 6 {
			continue // 跳过打印6 从7往下执行
		}
		fmt.Println(i)
	}
	fmt.Printf("contiune 跳过循环语句")
}

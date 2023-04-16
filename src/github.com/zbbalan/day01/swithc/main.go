package main

import (
	"fmt"
)

// switch判断

func main() {
	var n = 2
	if n == 1 {
		fmt.Println("A")
	} else if n == 2 {
		fmt.Println("B")
	} else if n == 3 {
		fmt.Println("C")
	}
	//switch 简化上面判断
	switch n := 3; n {
	case 1:
		fmt.Println("A")
	case 2:
		fmt.Println("B")
	case 3:
		fmt.Println("C")
	default:
		fmt.Println("无效数字")
	}
	switch m := 7; m {
	case 1, 3, 5, 7, 9:
		fmt.Printf("奇数") // 只要有一个能对上号就可以输出
	case 2, 4, 6, 8:
		fmt.Printf("偶数")
	default:
		fmt.Println(n)
	}
}

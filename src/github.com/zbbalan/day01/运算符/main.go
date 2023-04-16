package main

import (
	"fmt"
)

func main() {
	age := 26
	// && 并且
	if age > 18 && age < 60 {
		fmt.Printf("未退休接着上班")
	} else {
		fmt.Printf("退休啦")
	}
	// 或者||
	if age < 18 || age > 60 {
		fmt.Printf("未退休接着上班")
	} else {
		fmt.Printf("苦逼未退休接着上班")
	}
	// 位运算符：针对的是二进制运算
	//	5的二进制是：101
	//	2的二进制是：10

}

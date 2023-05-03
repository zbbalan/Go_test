package main

import (
	"fmt"
)

func main() {
	// copy 切片
	a1 := []int{1, 2, 3, 4, 5, 6}
	a2 := a1
	var a3 = make([]int, 6, 6) // 先定义a3 数组长度和容量
	copy(a3, a1)               // 将 a1的值拷贝给a3
	fmt.Println(a1, a2, a3)
	a1[0] = 1000
	fmt.Println(a1, a2, a3)
	var a4 = make([]int, 5, 15)
	for i := 0; i < 5; i++ {
		copy(a4, a3)
	}
	fmt.Println(a4)

}

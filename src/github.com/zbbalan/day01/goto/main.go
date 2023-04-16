package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j == 4 {
				goto LABLE // 如果 循环到4 就跳到标签LABLE。不往下执行

			}

		}
	}
LABLE:
	fmt.Println("over")
}

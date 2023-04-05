package main

import (
	"fmt"
)

// if 条件判断
func main() {
	age := 19
	if age > 18 {
		fmt.Println("大于18岁")
	} else {
		fmt.Println("小于18岁")
	}

	// 多条件判断
	if age > 35 {
		fmt.Println("大于35")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("未成年")
	}
}

package main

import (
	"fmt"
)

func main() {
	// map
	var s1 map[string]int
	s1 = make(map[string]int, 10) //初始化 在开始的时候估算好map容量，避免在程序运行期间动态扩容
	s1["ZBB"] = 6666
	s1["Alan"] = 7777
	fmt.Println(s1)
	fmt.Println(s1["ZBB"])
	// map 遍历
	for k, v := range s1 {
		fmt.Println(k, v)
	}
	// 只遍历 value
	for _, v := range s1 {
		fmt.Println(v)
	}
	// 删除 delete
	delete(s1, "Alan")
	fmt.Println(s1)

	// map 和slice的组合
	// 元素类型为map的切片
	var s2 = make([]map[int]string, 10, 10) // 定义
	s2[0] = make(map[int]string, 1)         // 初始化 s2 的值为key 是int 类型 value 是string 类型长度为1
	s2[0][10] = "北京"                        // 赋值
	fmt.Println(s2)
	// 值为切片类型的map
	var s3 = make(map[string][]int, 10)
	s3["中国"] = []int{6, 2, 4}
	fmt.Println(s3)
}

package main

import (
	"fmt"
)

func main() {

	//  切片的定义
	var s1 []int    // 定义一个int类型的切片
	var s2 []string // 定义一个string 类型的切片
	fmt.Println(s1, s2)
	// 初始化
	s1 = []int{1, 2, 3, 4}
	s2 = []string{"沙河", "昌平", "北京"}
	fmt.Println(s1, s2)
	// 长度和容量 len 长度 ，cap 容量
	fmt.Printf("len(s1):%d cap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))
	// 由数组得到切片
	a1 := [...]int{1, 6, 7, 9, 7, 6, 7, 8, 9}
	s3 := a1[0:4]   // 将 a1 的0---4 位赋值给s3 // 也是基于一个数组切割 左包含右不包含
	fmt.Println(s3) // 输出结果为1679
	s4 := a1[2:5]
	fmt.Println(s4)
	// cap 容量是原始切片的容量
	fmt.Printf("len(s4):%d cap(s4):%d\n", len(s4), cap(s4))
	// 切片在切片
	s5 := a1[:6] // 从第0 ----6位 输出结果 1 6 7 9 7 6
	s6 := s5[3:]
	fmt.Println(s6) // 将s6 在切片 从第三位切片到结束 输出结果 9 7 6
	// 现在长度是s6 切片的长度 ，cap容量是s5 的容量为6
	fmt.Printf("len(s6):%d cap(s6):%d\n", len(s6), cap(s6))
	a1[7] = 6666 // 更改底层数组值将第七位 值改为 66666
	fmt.Println(a1)

}

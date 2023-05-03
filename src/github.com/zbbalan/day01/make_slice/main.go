package main

import (
	"fmt"
)

func main() {
	// make() 函数创造切片
	s1 := make([]int, 5, 10) // 5 代表长度 10 代表容量  如果不写容量 容量与长度一致
	fmt.Printf("s1=%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	// 切片的赋值拷贝
	s2 := [...]int{4, 5, 6, 7, 8, 9}
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))
	s3 := s2
	fmt.Printf("len(s3):%d cap(s3):%d\n", len(s3), cap(s3))
	s3[2] = 66666
	fmt.Println(s2, s3)
	// 切片的遍历
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
}

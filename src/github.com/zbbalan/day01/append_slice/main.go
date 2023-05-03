package main

import (
	"fmt"
)

func main() {
	// append 切片的扩容追加元素
	s1 := []string{"沙河", "北京", "昌平"}
	fmt.Println(s1)
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "中国") // 调用append函数必须指明原始切片变量接受返回值
	fmt.Println(s1)
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	s2 := []string{"美国", "天津", "日本", "新加坡"}
	s1 = append(s1, s2...) // 将一个切片加入另外一个切片内 ... 代表拆开
	fmt.Println(s1)
	fmt.Printf("s1:%v len(s1):%d cap(s1):%d\n", s1, len(s1), cap(s1))
	// 将索引为0的删除  删除沙河
	s1 = append(s1[:0], s1[1:]...)
	fmt.Println(s1)
	fmt.Printf("cap(s1):%d\n", cap(s1))
	var a = make([]int, 5, 10)
	for i := 0; i < 10; i++ {
		a = append(a, i) // append 追加
	}
	fmt.Println(a) // 输出结果为[0 0 0 0 0 0 1 2 3 4 5 6 7 8 9]
}

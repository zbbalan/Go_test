package main

import (
	"fmt"
)

func main() {
	// 数组是存放元素的容器 必须指定元素的类型和长度
	var a1 [3]bool
	var a2 [4]bool
	fmt.Printf("a1:%T,a2:%T\n", a1, a2)
	// 数组初始话，，如果不初始化，默认元素都是零值（布尔值：false，整型和浮点型都是0，字符串："")
	fmt.Println(a1, a2)
	// 初始化第一种方式
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 初始化第二种方式,根据初始值自动推断长度
	a10 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}
	fmt.Println(a10)
	// 初始化第三种方式：根据索引来初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组遍历
	citys := [...]string{"北京", "天津", "河北", "山东", "浙江"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//多维数组
	var a12 [3][2]int // 多维数。先定义，在赋值 定义为3个数组，长度为2
	a12 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{4, 5},
	}
	fmt.Println(a12)
	// 多维数组遍历
	for _, v1 := range a12 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	// 求数组[1,3,5,7,8]数组的和
	B1 := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for _, C1 := range B1 {
		sum = sum + C1
	}
	fmt.Println(sum)

}

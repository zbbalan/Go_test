package main

import (
	"fmt"
	"strings"
)

// go 语言字符串只能用双引号 字符使用单引号,字符只能使用一个
func main() {
	string1 := "hello,world"
	string2 := 'h'
	string3 := '1'
	//path =: "'C:\\Users\\Horde\\Desktop\\Go_test\\src\\github.com'"
	fmt.Printf("%T,%s\n", string1, string1)
	fmt.Printf("%T,%s\n", string2, string2)
	fmt.Printf("%T,%s\n", string3, string3)
	//fmt.Printf(path)

	//定义多行字符串 使用反引号
	s2 := `
	   alan
	   jack
	   black
	   jackCheng 
	
	`
	fmt.Printf(s2)
	fmt.Println(len(string1)) // 打印字符串长度
	// 字符串拼接
	name := "理想"
	car := "汽车"
	fmt.Printf(name + car)
	// 单独一种写法，赋值拼接，在打印
	licar := fmt.Sprintf("%s%s", name, car)
	fmt.Println(licar)
	// 包含 licar这个变量里是否包含汽车这个词
	fmt.Println(strings.Contains(licar, "汽车")) // 返回值为true
}

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "hello,沙河"
	// 依次拿到字符串的字符
	//  判断当前字符是否是汉字
	// 把汉字出现的次数累加
	var count int
	for _, c := range s1 {
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	fmt.Println(count)

	// 单词出现的次数 how do you do
	s2 := "how go you do"
	s3 := strings.Split(s2, "")    // 把字符串按照空格切割得到切片
	m1 := make(map[string]int, 10) // 遍历切片存储到一个map
	for _, w := range s3 {
		// 如果map中不存在w这个key值那么次数为1
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++ // 如果存在key值那么次数+1
		}
	}
	for key, value := range m1 {
		fmt.Println(key, value)
	}
}

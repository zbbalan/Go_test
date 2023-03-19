// 程序入口
package main

//导入语句
import "fmt"

// 函数外只能放置标识符（变量/常量/函数/类型）的申明
// 程序入口函数
func main() {
	fmt.Println("helloworld")
}

// go run main.go 编译运行
// go build 执行二进制文件
// go install 先编译 在拷贝到bin目录下

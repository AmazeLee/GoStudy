package main

import "fmt"

func main0101() {

	var a byte = 'a'
	var b string = "a" // 'a' '\0'

	// 换行\n  \\表示一个\  一般用于文件操作
	var c string = "hello\nworld"
	// %s 遇到\0停止
	fmt.Println("%S", b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func main() {

	var str string = "hello world"
	// 在go语言中一个汉字算作3个字符 为了和linux同一处理
	var str1 string = "深圳福田"
	// 计算字符串个数
	num := len(str)
	num1 := len(str1)

	fmt.Println(num)
	fmt.Println(num1)
}

package day01

import "fmt"

func main0901() {

	// 布尔类型 默认值为false
	//var a bool
	//a = true
	//fmt.Println(a)

	// 自动推导类型创建bool类型变量
	a := false

	// %T是一个占位符 表示输出一个变量对应的数据类型
	fmt.Printf("%T", a)
}

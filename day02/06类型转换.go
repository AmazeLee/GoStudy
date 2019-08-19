package main

import "fmt"

func main0601() {

	a := 10
	b := 3.14

	// 将不同类型转成相同类型进行计算操作
	// 类型转换格式 数据类型（变量）
	// c := float64(a)*b

	// 将浮点型转为整型数据 保留浮点型整数部分 舍弃小数部分 不会进行四舍五入
	c := a * int(b)
	fmt.Println(c)
}

func main0602() {

	// 虽然都是整型，但是不允许转换，只有类型匹配的数据才能进行计算
	var a int32 = 10
	var b int64 = 20

	// 在go语言中 习惯将低类型转换为高类型
	c := int64(a) + b
	fmt.Println(c)
	fmt.Printf("%T", c)
}

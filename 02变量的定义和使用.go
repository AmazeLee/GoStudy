package main

import (
	"fmt"
	"math"
)

func main0201(){
	// 变量的定义和使用
	// var 变量名 数据类型 = 值
	// int 表示整型数据
	var sun int = 50

	fmt.Println(sun)
}

func main0202()  {

	// 变量的声明 如果没有赋值 这里默认为0
	var sun int

	// 为变量赋值
	sun = 50

	fmt.Println(sun)
}

func main0203()  {

	// float64 浮点型数据
	var value float64 = 2

	// var sum = value * value

	// 可以使用系统提供的包
	var sum = math.Pow(value,2)

	fmt.Println(sum)
}
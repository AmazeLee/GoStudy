package day01

import "fmt"

func main0301() {

	// 自动推导类型
	a := 10          // int 整型
	b := 20.234      // float64 浮点型
	c := "人生苦短，必须够浪" // string 字符串型

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a)
}

func main0302() {

	// 自动推导类型
	a := 10 // int 整型
	b := 20
	c := a
	a = b
	b = c

	fmt.Println(a)
	fmt.Println(b)
}

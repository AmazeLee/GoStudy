package day01

import "fmt"

func main0400() {

	// 多重赋值
	a, b, c := 10, 3.14, "人生苦短，必须够浪"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func main0401() {

	var a int = 10
	var b int = 20

	// 在一个作用域范围内变量名不能重复
	// var a float64 = 3.14

	// 如果在多重赋值时有新定义的变量 可以使用自动推导类型
	a, b, c, d := 110, 120, "你好", "朋友"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

}

func main0402() {

	// _表示匿名变量 不接收数据
	_, c, d := 120, "你好", "朋友"

	fmt.Println(c)
	fmt.Println(d)

}

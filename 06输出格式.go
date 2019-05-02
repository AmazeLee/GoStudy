package main

import "fmt"

func main() {

	fmt.Println("hello world")
	fmt.Print("深圳福田体育公园")

	a := 10
	b := 3.14159
	// %d是一个占位符 表示输出一个整型数据
	// %f是一个占位符 表示输出一个浮点型 保留6位小数 因为浮点型数据不是精准的数据 六位是有效的
	// %.2f保留小数位数为两位，会对第三位小数进行四舍五入
	fmt.Printf("%d %f %.2f\n", a, b, b)

	c := "你瞅啥";
	// %s是一个占位符 表示输出一个字符串类型
	fmt.Printf("%s", c)
}

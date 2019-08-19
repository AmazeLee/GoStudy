package main

import "fmt"

func main0101() {

	a := 10
	b := 20

	// 匿名内部函数
	// f 是函数类型对应的变量
	f := func(a int, b int) {
		fmt.Println(a + b)
	}

	f(a, b)
}

func main() {

	a := 10
	b := 20

	/*func(a int,b int)int{
		return a + b
	}(a,b)*/

	f := func(a int, b int) int {
		return a + b
	}

	v := f(a, b)
	fmt.Printf("%T", f)
	fmt.Printf("%T", v)
}

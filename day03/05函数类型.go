package main

import "fmt"

func test01() {

	fmt.Println("瓜娃子")
}

func test02(a int, b int) {
	fmt.Println(a + b)
}

// type 可以定义函数类型
// type 可以为已存在类型起别名
type FUNCTYPE func()
type FUNCTEST func(int, int)

func main() {

	// 定义函数类型变量
	var f FUNCTYPE
	f = test01
	f()

	var f1 FUNCTEST
	f1 = test02
	f1(1, 2)

	// 函数调用
	// 如果使用Print打印函数名是一个地址
	// 函数名本身就是一个指针类型数据 在内存中代码区进行存储
	fmt.Println(test01)

	// 自动类型推导创建函数类型、
	f3 := test02
	f3(2, 5)

}

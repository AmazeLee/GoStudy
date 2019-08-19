package main

import "fmt"

func main0201() {

	// 数组定义 var 数组名 [元素个数]数据类型

	// 切片定义 var 切片名 []数据类型
	/*var s []int
	fmt.Println(s)*/

	// 自动推导类型创建切片 make([]数据类型，5)
	s := make([]int, 5)
	// 通过下标为切片赋值
	s[0] = 123
	s[1] = 32
	s[2] = 56
	s[3] = 234

	// 通过append添加切片信息
	s = append(s, 678)

	fmt.Println(s)

	// 通过len查看切片长度
	fmt.Println(len(s))

	// 遍历
	for i, v := range s {
		fmt.Println(i, v)
	}
}

func main() {

	// 不写元素个数叫切片 必须写元素个数的叫数组
	var s []int = []int{1, 2, 3, 4, 5}

	// 遍历
	for i, v := range s {
		fmt.Println(i, v)
	}

	fmt.Println("长度：", len(s))
	fmt.Println("容量：", cap(s))
}

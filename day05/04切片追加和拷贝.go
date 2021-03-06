package main

import "fmt"

func main0401() {

	var s []int
	// s[0] = 123
	// 在使用append添加数据时 切片的地址可能会发生变化 如果容量扩充导致输出存储溢出
	// 切片会自动寻找新的空间存储数据，同时也会将之前的数据进行释放，使用append添加数据
	s = append(s, 1, 2, 3, 4, 5)

	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Printf("%p\n", s)

	s = append(s, 1, 2, 3, 4, 5)

	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Printf("%p\n", s)

}

func main0402() {

	s1 := []int{1, 2, 3, 4, 5}

	s2 := make([]int, 5)

	// 将s1中的数据拷贝到s2中，s2需要有足够的容量
	// 使用拷贝操作后，s1,s2是两个独立的空间 不会相互影响
	copy(s2, s1)

	fmt.Print(s2)
}

func main() {

	s1 := []int{1, 2, 3, 4, 5}

	s2 := make([]int, 5)

	// 如果想要拷贝切片中片段，需要使用切片中的截取
	copy(s2, s1[1:4])

	fmt.Print(s2)
}

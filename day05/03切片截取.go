package main

import "fmt"

func main0301() {

	s := []int{1, 2, 3, 4, 5}

	// 切片名[起始下标:]
	// slice := s[2:]

	// 切片名[:结束位置] 不包含结束位置
	//slice := s[:2]

	// 切片名[起始位置:结束位置]
	// slice := s[1:2]

	// 切片名[起始位置:结束位置:容量]
	slice := s[1:3:4]

	fmt.Println(slice)
	fmt.Println(cap(slice))
}

func main() {

	s := []int{1, 2, 3, 4, 5}

	slice := s[1:3]

	fmt.Printf("%p\n", s)
	fmt.Printf("%p\n", slice)
}

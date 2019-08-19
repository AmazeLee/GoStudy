package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5}

	// 切片名[起始下标:]
	// slice := s[2:]

	slice := s[:2]

	fmt.Println(slice)
}

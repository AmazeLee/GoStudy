package main

import "fmt"

func main() {

	test(1, 2, 33)
	test(1, 2)
}

func test(args ...int) {

	fmt.Println(args)
}

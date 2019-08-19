package main

import "fmt"

func main0401() {
	c := sub(5, 2)
	fmt.Println(c)
}

func sub(a int, b int) int {

	sum := a - b
	return sum
}

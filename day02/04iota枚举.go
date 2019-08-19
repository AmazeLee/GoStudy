package main

import "fmt"

/**
iota枚举
*/
func main0401() {

	const (
		a = iota
		b = iota
		c = iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

}

func main0402() {

	/*const(
		a = iota
		b
		c
	)*/

	const (
		a    = 10
		b, c = iota, iota
		d, e
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

package main

import "fmt"

func test(s []int) {

	fmt.Printf("%p\n", s)
}

func main() {

	s := []int{1, 2, 3, 4, 5}
	test(s)
	fmt.Printf("%p\n", s)

	lee := []int{9, 1, 34, 12, 10, 9}

	// 切片作为函数参数是地址传递 形参可以改变实参的值
	BubbleSort(lee)
	fmt.Println(lee)
}

// 冒泡排序
func BubbleSort(s []int) {

	for i := 0; i < len(s)-1; i++ {

		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {

				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}

}

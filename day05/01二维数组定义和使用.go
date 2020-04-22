package main

import "fmt"

func main() {

	// var 数组名 [元素个数] 数据类型  -- 一维数组
	// var 数组名 [行个数] [列个数] 数据类型 -- 二维数组
	var arr [3][4]int

	arr[1][3] = 4

	arr[2][1] = 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

}

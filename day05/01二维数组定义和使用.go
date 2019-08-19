package main

import "fmt"

func main() {

	var arr [3][4]int

	arr[1][3] = 4

	arr[2][1] = 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Println(arr[i][j])
		}
	}

}

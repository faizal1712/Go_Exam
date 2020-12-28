package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}

	matrix := make([][]int, len(arr[0]))
	for i := 0; i < len(arr[0]); i++ {
		matrix[i] = make([]int, len(arr))
	}

	for i := 0; i < len(arr[0]); i++ {
		for j := 0; j < len(arr); j++ {
			// fmt.Println(arr[j][i])
			matrix[i][j] = arr[j][i]
		}
	}
	fmt.Println(matrix)
}

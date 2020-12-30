package main

import "fmt"

func main() {
	arr := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{10, 11, 12},
	}
	fmt.Println(transposeMatrix(arr))
}

func transposeMatrix(arr [][]int) [][]int {
	arr0_length := len(arr[0])
	arr_length := len(arr)
	matrix := make([][]int, arr0_length)
	for i := 0; i < arr0_length; i++ {
		matrix[i] = make([]int, arr_length)
	}

	for i := 0; i < arr0_length; i++ {
		for j := 0; j < arr_length; j++ {
			matrix[i][j] = arr[j][i]
		}
	}
	return matrix
}

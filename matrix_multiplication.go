package main

import "fmt"

func main() {
	a := [][]int{
		{1, 2},
		{3, 4},
		{2, 1},
	}
	b := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	a0_length := len(a[0])
	b_length := len(b)
	if a0_length != b_length {
		fmt.Println("For a*c multiplication, given matrices should be in form of a*b and b*c")
		return
	}
	fmt.Println(multiplicationMatrix(a, b))
}

func multiplicationMatrix(a, b [][]int) [][]int {
	a_length := len(a)
	a0_length := len(a[0])
	b0_length := len(b[0])
	matrix := make([][]int, a_length)
	for i := 0; i < a_length; i++ {
		matrix[i] = make([]int, b0_length)
	}
	for i := 0; i < a_length; i++ {
		for j := 0; j < b0_length; j++ {
			for k := 0; k < a0_length; k++ {
				matrix[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return matrix
}

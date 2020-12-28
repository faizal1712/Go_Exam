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
	if len(a[0]) != len(b) {
		fmt.Println("For a*c multiplication, given matrices should be in form of a*b and b*c")
		return
	}
	matrix := make([][]int, len(a))
	for i := 0; i < len(a); i++ {
		matrix[i] = make([]int, len(b[0]))
	}
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(a[0]); k++ {
				matrix[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	fmt.Println(matrix)
}

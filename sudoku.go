package main

import "fmt"

func main() {
	var arr = [][]int{
		{8, 2, 7, 1, 5, 4, 3, 9, 6},
		{9, 6, 5, 3, 2, 7, 1, 4, 8},
		{3, 4, 1, 6, 8, 9, 7, 5, 2},
		{5, 9, 3, 4, 6, 8, 2, 7, 1},
		{4, 7, 2, 5, 1, 3, 6, 8, 9},
		{6, 1, 8, 9, 7, 2, 4, 3, 5},
		{7, 8, 6, 2, 3, 5, 9, 1, 4},
		{1, 5, 4, 7, 9, 6, 8, 2, 3},
		{2, 3, 9, 8, 4, 1, 5, 6, 7},
	}

	if !rowscheck(arr) || !columnscheck(arr) || !submatrixcheck(arr) {
		fmt.Println("This solution is wrong ! ")
		return
	}
	fmt.Println("This solution is true !")
}

func checkLine(arr []int) bool {
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	for i := 1; i < 10; i++ {
		if dict[i] != 1 {
			return false
		}
	}
	return true
}

func rowscheck(arr [][]int) bool {
	for i := 0; i < 9; i++ {
		if !checkLine(arr[i]) {
			return false
		}
	}
	return true
}

func columnscheck(arr [][]int) bool {
	for i := 0; i < 9; i++ {
		temp := make([]int, 9)
		for j := 0; j < 9; j++ {
			temp[j] = arr[j][i]
		}
		if !checkLine(temp) {
			return false
		}
	}
	return true
}

func submatrixcheck(arr [][]int) bool {
	temp := make([]int, 9)
	i, j := 0, 0
	for f := 0; f < 9; f++ {
		j = 3 * (f % 3)
		temp[0] = arr[i+0][j+0]
		temp[1] = arr[i+0][j+1]
		temp[2] = arr[i+0][j+2]
		temp[3] = arr[i+1][j+0]
		temp[4] = arr[i+1][j+1]
		temp[5] = arr[i+1][j+2]
		temp[6] = arr[i+2][j+0]
		temp[7] = arr[i+2][j+1]
		temp[8] = arr[i+2][j+2]
		if !checkLine(temp) {
			return false
		}
		if f%3 == 2 {
			i = i + 3
		}
	}
	return true
}

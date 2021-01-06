package main

import (
	"fmt"
	"time"
)

type coords struct {
	x int
	y int
}

const (
	size       = 9
	small_size = 3
)

func main() {
	now := time.Now()
	arr := map[int][]int{
		0: []int{0, 7, 0, 0, 0, 0, 0, 0, 9},
		1: []int{5, 1, 0, 4, 2, 0, 6, 0, 0},
		2: []int{0, 8, 0, 3, 0, 0, 7, 0, 0},
		3: []int{0, 0, 8, 0, 0, 1, 3, 7, 0},
		4: []int{0, 2, 3, 0, 8, 0, 0, 4, 0},
		5: []int{4, 0, 0, 9, 0, 0, 1, 0, 0},
		6: []int{9, 6, 2, 8, 0, 0, 0, 3, 0},
		7: []int{0, 0, 0, 0, 1, 0, 4, 0, 0},
		8: []int{7, 0, 0, 2, 0, 3, 0, 9, 6},
	}
	// fmt.Println(arr[1][1])
	var loc []coords
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if arr[i][j] != 0 {
				place := coords{i, j}
				loc = append(loc, place)
			}
		}
	}
	arr, flag := solvesudoku(0, 0, arr, loc)
	fmt.Println(time.Since(now))
	if flag && checksudoku(arr, true) {
		for i := 0; i < size; i++ {
			fmt.Println(arr[i])
		}
	} else {
		fmt.Println("This sudoku can't be solved ! ")
	}
}

func solvesudoku(x, y int, arr map[int][]int, loc []coords) (map[int][]int, bool) {
	for point := range loc {
		if loc[point].x == x && loc[point].y == y {
			flag := true
			if y != 8 {
				arr, flag = solvesudoku(x, y+1, arr, loc)
			} else if x != 8 {
				arr, flag = solvesudoku(x+1, 0, arr, loc)
			}
			return arr, flag
		}
	}
	if arr[x][y] == 0 || !checksudoku(arr, false) {
		arr[x][y] = arr[x][y] + 1
	}
	flag := false
	for i := arr[x][y]; i <= size && flag == false; i++ {
		if !checksudoku(arr, false) {
			arr[x][y] = arr[x][y] + 1
		} else {
			flag = true
			if y != 8 {
				arr, flag = solvesudoku(x, y+1, arr, loc)
			} else if x != 8 {
				arr, flag = solvesudoku(x+1, 0, arr, loc)
			}
			if !flag {
				arr[x][y] = arr[x][y] + 1
			}
		}
		// fmt.Println(x, y, checksudoku(arr, false), arr, flag)
	}
	if arr[x][y] == 10 || flag == false {
		arr[x][y] = 0
		return arr, false
	}
	return arr, checksudoku(arr, false)
}

func checksudoku(arr map[int][]int, issolved bool) bool {
	if !rowscheck(arr, issolved) || !columnscheck(arr, issolved) || !submatrixcheck(arr, issolved) {
		return false
	}
	return true
}

func checkLine(arr []int, issolved bool) bool {
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	for i := 1; i <= size; i++ {
		if issolved && dict[i] != 1 {
			return false
		} else if !issolved && dict[i] > 1 {
			return false
		}
	}
	return true
}

func rowscheck(arr map[int][]int, issolved bool) bool {
	for i := 0; i < size; i++ {
		if !checkLine(arr[i], issolved) {
			return false
		}
	}
	return true
}

func columnscheck(arr map[int][]int, issolved bool) bool {
	for i := 0; i < size; i++ {
		temp := make([]int, size)
		for j := 0; j < size; j++ {
			temp[j] = arr[j][i]
		}
		if !checkLine(temp, issolved) {
			return false
		}
	}
	return true
}

func submatrixcheck(arr map[int][]int, issolved bool) bool {
	temp := make([]int, size)
	i, j := 0, 0
	for f := 0; f < size; f++ {
		j = small_size * (f % small_size)
		temp[0] = arr[i+0][j+0]
		temp[1] = arr[i+0][j+1]
		temp[2] = arr[i+0][j+2]
		temp[3] = arr[i+1][j+0]
		temp[4] = arr[i+1][j+1]
		temp[5] = arr[i+1][j+2]
		temp[6] = arr[i+2][j+0]
		temp[7] = arr[i+2][j+1]
		temp[8] = arr[i+2][j+2]
		if !checkLine(temp, issolved) {
			return false
		}
		if f%small_size == (small_size - 1) {
			i = i + small_size
		}
	}
	return true
}

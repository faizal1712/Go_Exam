package main

import "fmt"

func main() {
	var arr = [9][9]int{
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
	// for i := 0; i < 9; i++ {
	// 	// for j := 0; j < 9; j++ {
	// 	// 	fmt.Print("Enter arr[", i, "][", j, "] sudoku value: ")
	// 	// 	fmt.Scanf(&arr[i][j])
	// 	// }
	// 	fmt.Print("Enter value: ")
	// 	fmt.Scanf("%d %d %d %d %d %d %d %d %d\n", &arr[i][0], &arr[i][1], &arr[i][2], &arr[i][3], &arr[i][4], &arr[i][5], &arr[i][6], &arr[i][7], &arr[i][8])
	// }
	// fmt.Println("Entered sudoku is: ")
	// for i := 0; i < 9; i++ {
	// 	for j := 0; j < 9; j++ {
	// 		fmt.Print(arr[i][j], "	")
	// 	}
	// 	fmt.Println()
	// }
	for i := 0; i < 9; i++ {
		if !checkLine(arr[i]) {
			fmt.Println("This solution is wrong ! ")
			return
		}
	}
	for i := 0; i < 9; i++ {
		var temp [9]int
		for j := 0; j < 9; j++ {
			temp[j] = arr[j][i]
		}
		if !checkLine(temp) {
			fmt.Println("This solution is wrong ! ")
			return
		}
	}
	var temp [9]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			for k := 0; k < 3; k++ {
				for l := 0; l < 3; l++ {
					temp[3*k+l] = arr[3*i+k][3*j+l]
				}
			}
			if !checkLine(temp) {
				fmt.Println("This solution is wrong ! ")
				return
			}
		}
	}
	fmt.Println("This solution is true !")
}

func checkLine(arr [9]int) bool {
	// fmt.Println(arr)
	dict := make(map[int]int)
	for _, num := range arr {
		dict[num] = dict[num] + 1
	}
	// fmt.Println(dict)
	for i := 1; i < 10; i++ {
		if dict[i] != 1 {
			return false
		}
	}
	return true
}

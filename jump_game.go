package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{0, 3, 2, 1, 0, 1, 1, 1, 1}
	fmt.Println("Given array is:", arr)
	// dest := len(arr)
	jum := jump(arr)
	if jum != math.MaxInt64 && jum != math.MinInt64 {
		fmt.Println("Minimum jump required to reach destination is : ", jum, math.MaxInt64)
	} else {
		fmt.Println("It's not possible to reach till end !!")
	}
}

func jump(nums []int) int {
	ans := make([]int, len(nums))
	ans[0] = 0
	for i := 1; i < len(nums); i++ {
		ans[i] = math.MaxInt64
	}
	// fmt.Println(ans)
	// if nums[0] == 0 {
	// 	return ans[len(ans)-1]
	// }
	for i := 0; i < len(nums); i++ {
		for j := i + 1; (j <= i+nums[i]) && (j < len(nums)); j++ {
			ans[j] = int(math.Min(float64(ans[j]), float64(ans[i]+1)))
			// fmt.Println(float64(ans[j]), float64(ans[i]+1))
		}
	}
	return ans[len(ans)-1]
}

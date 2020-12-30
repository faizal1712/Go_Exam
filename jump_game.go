package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{3, 2, 1, 1, 1, 1}
	fmt.Println("Given array is:", arr)
	jum := jump(arr)
	if jum != math.MaxInt64 && jum != math.MinInt64 {
		fmt.Println("Minimum jump required to reach destination is : ", jum)
	} else {
		fmt.Println("It's not possible to reach till end !!")
	}
}

func jump(nums []int) int {
	length := len(nums)
	ans := make([]int, length)
	ans[0] = 0
	for i := 1; i < length; i++ {
		ans[i] = math.MaxInt64
	}
	if nums[0] == 0 {
		return ans[length-1]
	}
	for i := 0; i < length; i++ {
		for j := i + 1; (j <= i+nums[i]) && (j < length); j++ {
			ans[j] = int(math.Min(float64(ans[j]), float64(ans[i]+1)))
		}
	}
	return ans[length-1]
}

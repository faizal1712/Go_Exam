package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var k, n int
	fmt.Print("Enter value of n: ")
	fmt.Scanln(&n)
	fmt.Print("Enter value of k: ")
	fmt.Scanln(&k)
	str := ""
	for i := 1; i <= n; i++ {
		str += strconv.Itoa(i)
	}
	p := fact(len(str))
	arr := make([]string, p)
	count := 0
	perm(0, str, len(str)-1, arr, &count)
	sort.Strings(arr)
	fmt.Println(k, "value of", n, "permutation is", arr[k-1])
}

func fact(no int) int {
	if no == 0 || no == 1 {
		return 1
	}
	return no * fact(no-1)
}

func perm(left int, str string, right int, arr []string, count *int) {
	if left == right {
		arr[*count] = str
		*count++
	} else {
		for i := left; i <= right; i++ {
			runeStr := []rune(str)
			runeStr[i], runeStr[left] = runeStr[left], runeStr[i]
			perm(left+1, string(runeStr), right, arr, count)
		}
	}
	return
}

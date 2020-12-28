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
	// fmt.Println(n, k)
	str := ""
	for i := 1; i <= n; i++ {
		str += strconv.Itoa(i)
	}
	// fmt.Println(str)
	p := fact(len(str))
	// fmt.Println(p)
	arr := make([]string, p)
	// char := make([]string, n)
	count := 0
	perm(1, str, arr, &count)
	// fmt.Println(arr)
	sort.Strings(arr)
	fmt.Println(arr)
	fmt.Println(k, "value of", n, "permutation is", arr[k-1])
}

func fact(no int) int {
	// defer fmt.Println("Factorial of ", no, "is completed")
	if no == 0 || no == 1 {
		return 1
	}
	return no * fact(no-1)
}

func replace(a int, str string, arr []string, count *int) {
	temp := str
	arr[*count] = temp
	// fmt.Println(temp)
	temp = str[:len(str)-2]
	temp += string(str[len(str)-1])
	temp += string(str[len(str)-2])
	*count++
	arr[*count] = temp
	// fmt.Println(temp)
	*count++
}

func perm(a int, str string, arr []string, count *int) {
	if len(str)-a < 2 {
		return
	}
	if len(str)-a == 2 {
		replace(a, str, arr, count)
	}
	perm(a+1, str, arr, count)
	// fmt.Println("in", a)
	for i := a; i < len(str); i++ {
		temp := ""
		// fmt.Println("in for", i, str, *count)
		if *count == fact(len(str)) {
			return
		}
		temp = str[:a-1]
		// fmt.Println(temp)
		temp += string(str[i])
		// fmt.Println(temp)
		// fmt.Println(temp)
		temp += str[a:i]
		// fmt.Println(temp)
		temp += str[i+1:]
		temp += string(str[a-1])
		// fmt.Println(temp)
		if len(str)-a == 2 {
			replace(a, temp, arr, count)
		} else {
			perm(a, temp, arr, count)
		}
	}
	return
}

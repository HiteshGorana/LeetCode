package main

import "fmt"

func productExceptSelf(nums []int) []int {
	n := len(nums)
	m := make([]int, n)
	start, end := 1, 1
	for i := 0; i < n; i++ {
		m[i] = start
		start *= nums[i]
	}

	for i := n - 1; i >= 0; i-- {
		m[i] *= end
		end *= nums[i]
	}
	return m
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

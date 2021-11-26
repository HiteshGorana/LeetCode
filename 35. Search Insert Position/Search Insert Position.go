package main

import "fmt"

func searchInsert(nums []int, target int) int {
	start, end := 0, len(nums)-1
	for start <= end {
		m := (start + end) / 2

		if nums[m] < target {
			start = m + 1
		} else {
			end = m - 1
		}
	}
	return start
}

func main() {
	fmt.Println(searchInsert([]int{1, 3, 5, 6}, 7))
}

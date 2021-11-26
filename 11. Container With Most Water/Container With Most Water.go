package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
func maxArea(height []int) int {
	n := len(height)
	max, start, end := 0, 0, n-1
	if n == 2 {
		return min(height[0], height[1])
	}
	for start < end {
		width := end - start
		h := 0
		if height[start] < height[end] {
			h = height[start]
			start++
		} else {
			h = height[end]
			end--
		}
		tem := h * width
		if tem > max {
			max = tem
		}
	}
	return max
}
func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

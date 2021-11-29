package main

import "fmt"

func removeDuplicates(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	return len(m)
}

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
}

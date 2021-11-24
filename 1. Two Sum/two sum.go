package main

import "fmt"

// twoSum : Main Idea O(n)
// 1. Iterate in array for each element
// 2. find other half
// 3. if other half exists in map return else save it in map
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, n := range nums {
		another := target - n
		if _, ok := m[another]; ok {
			return []int{m[another], index}
		}
		m[nums[index]] = index
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

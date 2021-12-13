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

If input array is sorted then
    - Binary search
    - Two pointers

If asked for all permutations/subsets then
    - Backtracking

If given a tree then
    - DFS
    - BFS

If given a graph then
    - DFS
    - BFS

If given a linked list then
    - Two pointers

If recursion is banned then
    - Stack

If must solve in-place then
    - Swap corresponding values
    - Store one or more different values in the same pointer

If asked for maximum/minumum subarray/subset/options then
    - Dynamic programming

If asked for top/least K items then
    - Heap

If asked for common strings then
    - Map
    - Trie

Else
    - Map/Set for O(1) time & O(n) space
    - Sort input for O(nlogn) time and O(1) space
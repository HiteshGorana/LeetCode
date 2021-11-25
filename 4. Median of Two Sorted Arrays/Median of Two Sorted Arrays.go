package main

import (
	"sort"
)

func median(arr []int) float64 {
	n := len(arr)
	if n%2 == 0 {
		output := 0
		for _, number := range arr[n/2-1 : n/2+1] {
			output += number
		}
		return float64(output) / 2
	} else {
		return float64(arr[n/2])
	}
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	num := append(nums1, nums2...)
	sort.Ints(num)
	return median(num)
}

func main() {
	findMedianSortedArrays([]int{1, 3, 4}, []int{2})
}

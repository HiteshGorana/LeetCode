package main

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	KEY := 0
	VALUE := 0
	for key, value := range m {
		if VALUE < value {
			KEY = key
			VALUE = value
		}
	}
	return KEY
}

func main() {
	majorityElement([]int{3, 2, 3})
}

// Runtime: 27 ms, faster than 18.52% of Go online submissions for Majority Element.
// Memory Usage: 6.4 MB, less than 14.60% of Go online submissions for Majority Element.
// Next challenges:

package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	right, left, size := 0, 0, len(s)-1

	for left < size {

		left++
		fmt.Printf("right %c left %c \n", s[right], s[left])
	}
	return 0
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

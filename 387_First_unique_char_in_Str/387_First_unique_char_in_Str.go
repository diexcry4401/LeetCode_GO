package main

import (
	"fmt"
)

/*

GIVEN:
String of lowercase eng chars

NEED:
Find indx of first unique char in a string
If it doesn't exist return -1

"leetcode" - return 0

"aabbcddee" - return 4, s[4] = c


IDEAS:
Try to use two pointers, first will be a pointer to unique char,
second will iterate through the string and compare with value of first pointer

*/

func firstUniqChar(s string) int {
	ans := -1
	m := map[rune]int{}
	for _, val := range s {
		m[val]++
	}
	for i, val := range s {
		if m[val] == 1 {
			ans = i
			break
		}
	}
	return ans
}

func main() {
	ex1 := "llllo"

	fmt.Println(firstUniqChar(ex1))
}

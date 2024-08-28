package main

import "fmt"

func isPalindrome(s string) bool {
	m := []int{}
	count := 0
	for _, val := range s {
		if (int(val) >= 97 && int(val) <= 122) || (int(val) >= 48 && int(val) <= 57) {
			m = append(m, int(val))
			count++
		}
		if int(val) >= 65 && int(val) <= 90 {
			m = append(m, int(val)+32)
			count++
		}
	}
	start := 0
	end := count - 1
	for start < end {
		if m[start] != m[end] {
			return false
		}
		start++
		end--
	}
	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	a := "race a car"
	b := "0P"
	fmt.Println(isPalindrome(s))
	fmt.Println(isPalindrome(a))
	fmt.Println(isPalindrome(b))
}

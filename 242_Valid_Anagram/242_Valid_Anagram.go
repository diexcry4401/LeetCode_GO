package main

import "fmt"

func isAnagram(s string, t string) bool {
	chars := make(map[rune]int)
	for _, val := range s {
		chars[val]++
	}
	for _, val := range t {
		chars[val]--
	}
	for _, v := range chars {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "ac"
	t := "bb"
	fmt.Println(isAnagram(s, t))
}

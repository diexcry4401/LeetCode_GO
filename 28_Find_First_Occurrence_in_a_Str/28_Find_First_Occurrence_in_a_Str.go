package main

import "fmt"

/*

sadbutsad sad
ans = 0

ferrari arir
ans = -1

abc c
ans = 2

решение ниже работает за O(N*M)

сделать через срезы за O(N)
*/

func strStr(haystack string, needle string) int {
	ans := -1
	for i := range haystack {
		if len(haystack)-i >= len(needle) {
			if haystack[i:i+len(needle)] == needle {
				fmt.Printf("haystack[%d:%d] = %s, needle = %s\n", i, i+len(needle), haystack[i:i+len(needle)], needle)
				return i
			}
		} else {
			return ans
		}
	}
	return ans
}

func main() {
	a := "sasasadbutsad"
	b := "ferrari"
	c := "abc"
	a1 := "sad"
	b1 := "arir"
	c1 := "c"
	fmt.Println(strStr(a, a1))
	fmt.Println(strStr(b, b1))
	fmt.Println(strStr(c, c1))
}

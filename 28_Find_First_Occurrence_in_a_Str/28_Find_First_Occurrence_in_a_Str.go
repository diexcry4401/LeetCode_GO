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
	for i, _ := range haystack {
		fmt.Printf("haystack[i] = %d, needle[0] = %d, %d, %d\n", haystack[i], needle[0], len(haystack)-1-i, len(needle)-1)
		if len(haystack)-1-i >= len(needle)-1 {
			if haystack[i] == needle[0] {
				count := 0
				haystackInd := i
				for ind, _ := range needle {
					fmt.Println(haystack[haystackInd], haystackInd, needle[ind])
					if haystack[haystackInd] == needle[ind] {
						count++
						haystackInd++
					} else {
						fmt.Println("Break")
						break
					}
				}
				if count == len(needle) {
					return i
				}
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

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func reverse(x int) {
	var start int
	edge := 2147483647
	if x >= edge || x <= (edge*-1)-1 {
		fmt.Println("Out of range")
	}
	s := strings.Split(strconv.Itoa(x), "")
	if s[0] == "-" {
		start = 1
	} else {
		start = 0
	}
	fmt.Println(s)
	end := len(s) - 1
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
	fmt.Println(s)
}

func main() {
	num := -123124
	reverse(num)
}

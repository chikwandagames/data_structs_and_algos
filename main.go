package main

import (
	"fmt"
	"strconv"
)

func main() {
	r := reverse(1534236469)
	fmt.Println(r)
}

func reverse(x int) int {

	s := strconv.Itoa(x)
	var ns string

	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) == "-" {
			continue
		}
		// fmt.Println(string(s[i]))
		ns += string(s[i])
	}
	num, _ := strconv.Atoi(ns)
	// fmt.Println(ns)

	if num > 2147483647 || num < -2147483648 {
		return 0
	}
	if x < 0 {
		return 0 - num
	}

	return num
}

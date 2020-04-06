// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Control structure
package main

import (
	"fmt"
)

func normalFor() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("normalFor: sum is %d\n", sum)
}
func forAsWhile() {
	var i int = 0
	var sum int = 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Printf("forAsWhile: sum is %d\n", sum)
}
func forRange() {
	for i, char := range "abcdef" {
		fmt.Printf("Index %d is %c\n", i, char)
	}
}

func main() {
	normalFor()
	forAsWhile()
	forRange()
}

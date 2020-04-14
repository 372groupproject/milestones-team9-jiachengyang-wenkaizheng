// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// String and Pointer type
package main

import (
	"fmt"
)

func String() {
	var p string = "abc"
	var q *string = &p
	fmt.Println(p)
	if *q == p {
		fmt.Println("P and Q have the same value")
	}
	*q = "ABC"
	fmt.Println(*q)
	fmt.Println(p)
	p += "abc"
	fmt.Println(p)
	fmt.Println(len(p))
	p = "English中文日本語は"
	p2 := p[10]
	fmt.Printf("%c\n", p2)
	for i, c := range p {
		fmt.Printf("%d , %c\n", i, c)

	}
}
func main() {
	String()
}

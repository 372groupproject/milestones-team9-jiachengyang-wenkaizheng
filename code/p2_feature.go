// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// 6 types of control structure
package main

import (
	"fmt"
)

func type1() {
	var sum int = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf("Type1 sum is %d\n", sum)
}
func type2() {
	var i int = 0
	var sum int = 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Printf("Type2 sum is %d\n", sum)
}
func type3(variable int) {
	if variable%2 == 0 {
		fmt.Println("Type3 is even")
	} else {
		fmt.Println("Type3 is odd")
	}
}
func type4() int {
	defer func() { fmt.Println("three") }()
	defer func() { fmt.Println("two") }()
	defer func() { fmt.Println("one") }()

	return 1
}
func type5(i int) {
	switch i {
	case 0:
		fmt.Println("Type5 is 0")
	case 1:
		fmt.Println("Type5 is 1")
	default:
		fmt.Println("Type5 is default")
	}
}
func type6(variable int) {
	switch {
	case variable%2 == 0:
		fmt.Println("Type6 is even")
	default:
		fmt.Println("Type6 is odd")
	}
}
func main() {
	type1()
	type2()
	type3(1)
	type4()
	type5(0)
	type6(0)
}

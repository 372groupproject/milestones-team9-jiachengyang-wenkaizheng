// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Control structure
package main

import (
	"fmt"
)

func normalIf(variable int) {
	if variable%2 == 0 {
		fmt.Println("normalIf: is even")
	} else {
		fmt.Println("normalIf: is odd")
	}
}

func ifWithVariableDefine() {
	if x := 1; x%2 == 0 {
		fmt.Print(x)
		fmt.Println(" is even")
	} else {
		fmt.Print(x)
		fmt.Println(" is odd")
	}
}

func main() {
	normalIf(1)
	ifWithVariableDefine()
}

// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Control structure
package main

import (
	"fmt"
)

func normalSwitch(i int) {
	switch i {
	case 0:
		fmt.Println("normalSwitch: case 0")
	case 1:
		fmt.Println("normalSwitch:  case 1")
	default:
		fmt.Println("normalSwitch: default")
	}
}
func switchAsIfElseIF(variable int) {
	switch {
	case variable%2 == 0:
		fmt.Println("variable is divisible by 2")
	case variable%3 == 0:
		fmt.Println("variable is divisible by 3")
	case variable%5 == 0:
		fmt.Println("variable is divisible by 5")
	case variable%7 == 0:
		fmt.Println("variable is divisible by 7")
	default:
		fmt.Println("variable is not divisible by 2, 3, 5 or 7")
	}
}
func main() {
	normalSwitch(0)
	switchAsIfElseIF(0)
}

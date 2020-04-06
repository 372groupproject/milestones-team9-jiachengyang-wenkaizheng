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
		fmt.Println("normalSwitch: case 1")
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

func switchWithAList(variable int) {
	switch variable {
	case 2, 3, 5, 7:
		fmt.Println("variable is a prime less than 10")
	default:
		fmt.Println("variable is not a prime less than 10")
	}
}

func switchWithFallthrough(variable int) {
	switch variable {
	case 2:
		fmt.Println("variable is a prime less than 3")
		fallthrough
	case 3:
		fmt.Println("variable is a prime less than 4")
		fallthrough
	case 5:
		fmt.Println("variable is a prime less than 6")
		fallthrough
	case 7:
		fmt.Println("variable is a prime less than 10")
	default:
		fmt.Println("variable is not a prime less than 10")
	}
}

func main() {
	normalSwitch(0)
	switchAsIfElseIF(0)
	switchWithAList(2)
	switchWithFallthrough(2)
}

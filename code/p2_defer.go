// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Control structure
package main

import (
	"fmt"
)

func deferPrintln() int {
	defer func() { fmt.Println("three") }()
	defer func() { fmt.Println("two") }()
	defer func() { fmt.Println("one") }()

	return 1
}

func main() {
	deferPrintln()
}

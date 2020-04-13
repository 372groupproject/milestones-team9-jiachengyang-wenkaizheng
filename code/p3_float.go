// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Float and Float Pointer
package main

import (
	"fmt"
)

func Float32() {
	var p float32 = 0
	var q float32 = 1
	fmt.Println(p + q)
	if p != q {
		fmt.Println("They don't have the same value")
	}
	var p1 float32 = 100
	var p2 float32 = 10
	fmt.Printf("Before swap p1 is %f and p2 is %f\n", p1, p2)
	var p3 *float32 = &p2
	var p4 *float32 = &p1
	var tmp float32 = *p3
	*p3 = *p4
	*p4 = tmp
	fmt.Printf("After swap p1 is %f and p2 is %f\n", p1, p2)

}
func Float64() {
	var p float64 = 0
	var q float64 = 1
	fmt.Println(p + q)
	if p != q {
		fmt.Println("They don't have the same value")
	}
	var p1 float64 = 100
	var p2 float64 = 10
	fmt.Printf("Before swap p1 is %f and p2 is %f\n", p1, p2)
	var p3 *float64 = &p2
	var p4 *float64 = &p1
	var tmp float64 = *p3
	*p3 = *p4
	*p4 = tmp
	fmt.Printf("After swap p1 is %f and p2 is %f\n", p1, p2)

}
func main() {
	Float32()
	Float64()
}

// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Int and Int Pointer
package main

import (
	"fmt"
)

func Int() {
	var p int = 0
	var q int = 1
	fmt.Println(p + q)
	if p != q {
		fmt.Println("They don't have the same value")
	}
	var p1 int = 100
	var p2 int = 10
	fmt.Printf("Before swap p1 is %d and p2 is %d\n", p1, p2)
	var p3 *int = &p2
	var p4 *int = &p1
	var tmp int = *p3
	*p3 = *p4
	*p4 = tmp
	fmt.Printf("After swap p1 is %d and p2 is %d\n", p1, p2)

}
func Uint() {
	var p uint = 0
	var q uint = 1
	fmt.Println(p + q)
	if p != q {
		fmt.Println("They don't have the same value")
	}
	var p1 uint = 100
	var p2 uint = 10
	fmt.Printf("Before swap p1 is %d and p2 is %d\n", p1, p2)
	var p3 *uint = &p2
	var p4 *uint = &p1
	var tmp uint = *p3
	*p3 = *p4
	*p4 = tmp
	fmt.Printf("After swap p1 is %d and p2 is %d\n", p1, p2)

}
func Int32() {
	var p int32 = 0
	var q int32 = 1
	fmt.Println(p + q)
	if p != q {
		fmt.Println("They don't have the same value")
	}
	var p1 int32 = 100
	var p2 int32 = 10
	fmt.Printf("Before swap p1 is %d and p2 is %d\n", p1, p2)
	var p3 *int32 = &p2
	var p4 *int32 = &p1
	var tmp int32 = *p3
	*p3 = *p4
	*p4 = tmp
	fmt.Printf("After swap p1 is %d and p2 is %d\n", p1, p2)
}
func Uint32() {
	var p uint32 = 0
	var q uint32 = 1
	fmt.Println(p + q)
	if p != q {
		fmt.Println("They don't have the same value")
	}
	var p1 uint32 = 100
	var p2 uint32 = 10
	fmt.Printf("Before swap p1 is %d and p2 is %d\n", p1, p2)
	var p3 *uint32 = &p2
	var p4 *uint32 = &p1
	var tmp uint32 = *p3
	*p3 = *p4
	*p4 = tmp
	fmt.Printf("After swap p1 is %d and p2 is %d\n", p1, p2)
}
func Int64() {
	var p int64 = 0
	var q int64 = 1
	fmt.Println(p + q)
	if p != q {
		fmt.Println("They don't have the same value")
	}
	var p1 int64 = 100
	var p2 int64 = 10
	fmt.Printf("Before swap p1 is %d and p2 is %d\n", p1, p2)
	var p3 *int64 = &p2
	var p4 *int64 = &p1
	var tmp int64 = *p3
	*p3 = *p4
	*p4 = tmp
	fmt.Printf("After swap p1 is %d and p2 is %d\n", p1, p2)
}
func Uint64() {
	var p uint64 = 0
	var q uint64 = 1
	fmt.Println(p + q)
	if p != q {
		fmt.Println("They don't have the same value")
	}
	var p1 uint64 = 100
	var p2 uint64 = 10
	fmt.Printf("Before swap p1 is %d and p2 is %d\n", p1, p2)
	var p3 *uint64 = &p2
	var p4 *uint64 = &p1
	var tmp uint64 = *p3
	*p3 = *p4
	*p4 = tmp
	fmt.Printf("After swap p1 is %d and p2 is %d\n", p1, p2)
}

func main() {
	Int()
	Int32()
	Int64()
	Uint()
	Uint32()
	Uint64()

}

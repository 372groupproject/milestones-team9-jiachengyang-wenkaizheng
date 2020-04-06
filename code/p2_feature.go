// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Hello World in different thread
package main

import (
	//  "runtime"
	"fmt"
)

func Hello() {
	for i := 0; i < 10; i++ {
		fmt.Print("Hello ")
		if i == 9 {
			ch3 <- 0
			break
		}
		ch3 <- 0
		<-ch4

	}
	ch1 <- 0
}
func World() {
	<-ch3
	for i := 0; i < 10; i++ {
		fmt.Print("Wrold \n")
		if i == 9 {
			break
		}
		ch4 <- 0
		<-ch3

	}
	ch2 <- 0
}

var ch3 = make(chan int)
var ch4 = make(chan int)
var ch1 = make(chan int)
var ch2 = make(chan int)

func main() {

	go Hello()
	go World()
	<-ch1
	<-ch2

}

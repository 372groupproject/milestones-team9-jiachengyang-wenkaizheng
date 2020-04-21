// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// different main thread and sub thread
package main

import (
	//  "runtime"
	"fmt"
	"math"
)

func producer(ch, complete chan int) {
	for i := 2; i < 10000; i++ {
		ch <- i
	}
	close(ch)
	complete <- 1
}

func consumer(ch, complete chan int) {
	for {
		v, ok := <-ch
		if !ok {
			complete <- 1
			return
		}
		if isFact(v) {
			fmt.Println(v)
		}
	}

}

func isFact(n int) bool {
	bound := int(math.Ceil(math.Pow(float64(n), 0.5)))
	for i := 2; i <= bound; i++ {
		if n != i && n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go producer(ch1, ch2)
	go consumer(ch1, ch2)
	go consumer(ch1, ch2)
	<-ch2
	<-ch2
	<-ch2
	fmt.Println("All subthread finished work")

}

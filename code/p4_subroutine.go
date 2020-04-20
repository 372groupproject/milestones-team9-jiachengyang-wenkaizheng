// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// different main thread and sub thread
package main
import (
  //  "runtime"
    "fmt"
   
)

func Even(){
   for i:=0;i<10;i+=2{
	  fmt.Printf("Even thread is running with %d\n",i)
      if i == 8{
          ch3 <-0
          break;
      }
      ch3 <-0
      <- ch4
      
            
    }
   ch1 <- 0 
}
func Odd(){
   <- ch3
   for i:=1;i<10;i+=2{
      fmt.Printf("Odd thread is running with %d\n",i)
      if i == 9{
          break
      }
      ch4 <- 0
      <- ch3

            
    }
   ch2 <- 0 
}
var ch3 = make(chan int)
var ch4 = make(chan int)
var ch1 =  make(chan int)
var ch2 = make(chan int)
func main() {
  
  go Even()
  go Odd()
  <-ch1
  <- ch2
  fmt.Println("All subthread finished work")

}
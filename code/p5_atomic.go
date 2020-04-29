package main

import (
	"fmt"
    "sync/atomic"
)

type Student struct {
	Age   int32
	Name  string
	Grade int32
    Lock  chan int 
}
func MallocStudent(age ,grade int32,name string, l chan int ) *Student{
     return &Student{age,name,grade,l}
}
func final(s *Student){
    atomic.AddInt32(&sum, s.Grade)
    if atomic.CompareAndSwapInt32(&already, 0, s.Grade) {
        fmt.Printf("Swap alreay by %s\n",s.Name)
    }
    s.Lock <-0
    
}
var already int32 = 0
var sum int32 = 0
func main() {
   
    one:= make(chan int)
	two:= make(chan int)
	three :=make(chan int)
	s1 :=MallocStudent(22, 99,"abc",one)
	s2 :=MallocStudent(23, 100,"ABC",two)
	s3 :=MallocStudent(24, 101,"ABCD",three)
    go final(s1)
	go final(s2)
	go final(s3)
    <-one 
	<-two 
	<-three
    fmt.Println(sum)
    fmt.Println(already)
}
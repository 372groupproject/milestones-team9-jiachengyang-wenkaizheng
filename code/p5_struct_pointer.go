package main

import (
	"fmt"
)

type Student struct {
	Age   int32
	Name  string
	Grade int32
}
func MallocStudent(age ,grade int32,name string) *Student{
     return &Student{age,name,grade}
}
func change (s *Student,age ,grade int32,name string){
       s.Age = age
       s.Name =name
       s.Grade = grade
}
func main() {
   
	s1 :=MallocStudent(22, 99,"abc")
    fmt.Printf("Name %s, Grade %d, Age %d\n" ,s1.Name,s1.Grade,s1.Age)
    change(s1, 23,100,"ABC")
    fmt.Printf("Name %s, Grade %d, Age %d\n" ,s1.Name,s1.Grade,s1.Age)
}
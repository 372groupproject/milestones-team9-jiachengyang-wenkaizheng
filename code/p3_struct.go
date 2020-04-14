package main

import "fmt"

type Student struct {
	Age   int
	Name  string
	Grade int
}

func main() {
	s := Student{Age: 22, Name: "Derrick", Grade: 100}
	pointer := &s
	fmt.Printf("%d %s %d\n", s.Age, s.Name, s.Grade)
	fmt.Printf("%d %s %d\n", pointer.Age, pointer.Name, pointer.Grade)
}

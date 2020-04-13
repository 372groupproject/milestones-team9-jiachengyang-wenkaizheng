package main

import "fmt"

func arrayOperation() {
	var intArr [30]int
	for i := 0; i < 30; i++ {
		intArr[i] = i
	}
	for i := 0; i < 30; i++ {
		fmt.Println(intArr[i])
	}
	stringArr := [6]string{"zero", "one", "two", "three", "four", "five"}
	for i := 0; i < len(stringArr); i++ {
		fmt.Println(stringArr[i])
	}
}
func main() {
	arrayOperation()
}

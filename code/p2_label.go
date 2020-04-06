package main

import "fmt"

func main() {
	var a int = 0
Loop:
	for a < 10 {
		fmt.Println(a)
		a++
		if a == 8 {
			fmt.Println("Going to break")
			break Loop
		}
		goto Loop
	}
}

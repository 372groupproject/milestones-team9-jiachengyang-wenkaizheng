package main

import (
	
	"os"
	"log"
	"fmt"
)
func main(){
	file, err := os.Open("no exist")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(file)
}
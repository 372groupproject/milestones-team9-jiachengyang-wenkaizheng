// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Control structure 
package main
import (
	"fmt"
)

func type5(i int){
    switch i { 
        case 0: 
            fmt.Println("Type5 is 0") 
        case 1: 
            fmt.Println("Type5 is 1") 
        default: 
            fmt.Println("Type5 is default") 
    } 
}
func type6 (variable int){
    switch { 
        case variable % 2 == 0: 
            fmt.Println("Type6 is even")
		default: 
            fmt.Println("Typr6 is odd") 
    }  
}
func main() {
    type5(0)
    type6(0)
}
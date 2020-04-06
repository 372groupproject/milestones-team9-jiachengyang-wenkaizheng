// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Control structure 
package main
import (
	"fmt"
)
func type3(variable int){
     if variable % 2 ==0 {
         fmt.Println("Type3 is even")
     }else{
         fmt.Println("Type3 is odd")
     }
}

func main() {
	type3(1)
}
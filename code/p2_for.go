// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Control structure 
package main
import (
	"fmt"
)
func type1(){
    var sum int = 0
	for i:=0;i<10;i++{
		sum+=i
	}
    fmt.Printf("Type1 sum is %d\n",sum)
}
func type2(){
    var i int  = 0
    var sum int = 0
    for ;i<10;{
       sum +=i
       i++
    }
    fmt.Printf("Type2 sum is %d\n",sum)
}

func main() {
	type1()
    type2()
}
package main

import (
	"fmt"
)

func SwapByPointer(a,b *int){
	tmp := *a
	*a = *b
	*b = tmp
}
func main() {
	 var a int =10
	 var b int =15
	 a,b = b,a
	 fmt.Printf("a is %d, b is %d\n",a,b)
	 SwapByPointer(&a,&b)
	 fmt.Printf("a is %d, b is %d\n",a,b)
}
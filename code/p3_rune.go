package main

import (
	"fmt"
	"strconv"
)

func rune_operation() {
	// byte
	var p8 rune = 'は'
	fmt.Printf("%#U %+q %c\n", p8, p8, p8)
	fmt.Println(strconv.QuoteRuneToASCII(p8))
}
func main() {
	rune_operation()
}

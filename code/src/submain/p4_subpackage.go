package main
import (
	
	"subhelper"
	"fmt"
)
func main(){
	subhelper.Printer()
	fmt.Println(subhelper.Add(1,2))
	fmt.Println(subhelper.Sub(1,2))
	fmt.Println(subhelper.Mul(1,2))
	fmt.Println(subhelper.Div(1,2))
}
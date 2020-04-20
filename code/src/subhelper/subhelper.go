package subhelper
import (
     "fmt"
)
func Printer(){
	for i:=0;i<10;i++{
		fmt.Println("Sub funtion has been call")
	}
}
func Add(a,b int)int{
	return a+b
}
func Sub(a,b int)int{
	return a-b
}
func Mul(a,b int)int{
	return a*b
}
func Div(a,b int)int{
	return a/b
}
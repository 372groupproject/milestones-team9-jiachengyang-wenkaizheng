package main
import "fmt"
func main(){
	var a int = 0
	Loop: for a < 10{
		fmt.Println(a)
		a++
		goto Loop
	}
}
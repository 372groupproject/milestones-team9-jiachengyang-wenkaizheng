package main 
  
import "fmt"

func array_operation(){
	var arr [30]int
	for i:=0;i<30;i++{
		arr[i] = i
	}
	for i:=0;i<30;i++{
		fmt.Println(arr[i])
	}
	arrs:= [6]string{"zero", "one", "two", "three","four","five"}
    for i:=0;i<len(arrs);i++{
		fmt.Println(arrs[i])
	}
}
func main() { 
   array_operation()
} 
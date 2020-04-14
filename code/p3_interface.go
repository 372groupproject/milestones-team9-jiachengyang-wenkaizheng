package main
import "fmt"

type method interface{
	add() int
}

type adder struct{
	First int
	Second int
}
func (a *adder) add()int{
	return a.First + a.Second
}
func interact(m method){
	fmt.Println(m.add())
}
func main(){
	a := adder{0,1}
	interact(&a)

}
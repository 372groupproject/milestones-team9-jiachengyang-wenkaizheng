package main
import "fmt"
type child struct{
    Num int
}
type parent struct{
    Child [4]*child
}
func MallocParent()*parent{
	c:=&child{Num:0}
	c1:=&child{Num:1}
	c2:=&child{Num:2}
	c3:=&child{Num:3}
	arr:= [4]*child{c,c1,c2,c3}
	return &parent{Child:arr}
}

func main(){
	p:=MallocParent()
	for i:=0;i<4;i++{
	  fmt.Println(p.Child[i].Num)
	}
}
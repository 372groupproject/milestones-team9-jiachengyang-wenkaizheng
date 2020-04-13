package main

import "fmt"

func map_operation(){
	m := make(map[int]string)

    m[0] = "zero"
    m[1] = "one"
	m[2] = "two"
	m[3] = "three"
	m[4] = "four"
	m[5] = "five"
    fmt.Println("Direct Println:", m)

    v := m[0]
    fmt.Println("vvalue: ", v)

    fmt.Println("Length:", len(m))

    delete(m, 0)
	fmt.Println("Direct Println:", m)
	
	m[0] = "zero"

   for k,v :=range m{
	   fmt.Printf("The key is %d and the value is %s\n",k,v)
   }
   m1 := map[string]int{"zero": 0, "one": 1,"two":2,"three":3,"four":4,"five":5}
   for k,v :=range m1{
	fmt.Printf("The key is %s and the value is %d\n",k,v)
}

}
func main() {

    map_operation()
}
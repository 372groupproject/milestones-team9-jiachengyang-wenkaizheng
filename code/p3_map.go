package main

import "fmt"

func mapOperation() {
	m := make(map[rune]string)

	m['零'] = "zero"
	m['一'] = "one"
	m['二'] = "two"
	m['三'] = "three"
	m['四'] = "four"
	m['五'] = "five"
	fmt.Println("Direct Println:", m)

	v := m['零']
	fmt.Println("value of v: ", v)

	fmt.Println("Length:", len(m))

	delete(m, '零')
	fmt.Println("Direct Println:", m)

	m['零'] = "zero"

	for k, v := range m {
		fmt.Printf("The key is %c and the value is %s\n", k, v)
	}
	m1 := map[string]int{"zero": 0, "one": 1, "two": 2, "three": 3, "four": 4, "five": 5}
	for k, v := range m1 {
		fmt.Printf("The key is %s and the value is %d\n", k, v)
	}

}
func main() {

	mapOperation()
}

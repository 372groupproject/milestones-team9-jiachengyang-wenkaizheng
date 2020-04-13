package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func CountFile(m map[string]int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the name of file: ")
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	for {
		var key string
		n, err := fmt.Fscanf(file, "%s", &key)
		//fmt.Println(key)
		if err != nil {
			n++
			break
		}
		if value, exist := m[key]; exist {
			m[key] = value + 1
		} else {
			m[key] = 1
		}

	}

}
func main() {
	m := make(map[string]int)
	CountFile(m)
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i:=0;i<len(m);i++ {
		fmt.Printf("%s : %d\n", keys[i], m[keys[i]])
	}
}

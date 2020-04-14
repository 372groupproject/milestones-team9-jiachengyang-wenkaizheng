package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func output(m map[string]int) {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for i := 0; i < len(m); i++ {
		fmt.Printf("%s : %d\n", keys[i], m[keys[i]])
	}
}
func dealRedirect(line string, m map[string]int) {
	words := strings.Fields(line)
	for i := 0; i < len(words); i++ {
		key := words[i]
		if value, exist := m[key]; exist {
			m[key] = value + 1
		} else {
			m[key] = 1
		}

	}
}
func dealArgu(name string, m map[string]int) {
	file, err := os.Open(name)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		dealRedirect(line, m)
	}
}
func main() {
	m := make(map[string]int)
	if len(os.Args) >= 2 {
		dealArgu(os.Args[1], m)
		output(m)
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		dealRedirect(line, m)
		fmt.Println(line)
	}
	output(m)

}

package main

import (
    "io/ioutil"
    "encoding/json"
    "fmt"
    "os"
)
type network struct {
    Server      string   `json:"server"`
    Server_port    int     `json:"server_port"`
    Local_port    int     `json:"local_port"`
    Password     string   `json:"password"`
    Timeout      int      `json:"timeout"`
}

func file_get_contents(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}
 
func main() {
    fmt.Println("Hello World")
	var c network
	var content []byte
	var err error
	
	content, err = file_get_contents("./config.json")
	if err != nil {
		fmt.Println("open file error: "+ err.Error())
		return
	}
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}
	fmt.Println(c);
	
}

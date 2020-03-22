package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type network struct {
	Server     string `json:"server"`
	ServerPort int    `json:"server_port"`
	LocalPort  int    `json:"local_port"`
	Password   string `json:"password"`
	Timeout    int    `json:"timeout"`
}

var configPATH = "./config.json"

func readJson(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func main() {
	fmt.Println("Hello World")

	var c network

	content, err := readJson(configPATH)
	if err != nil {
		fmt.Println("Cannot open " + configPATH + " : " + err.Error())
		return
	}
	err = json.Unmarshal([]byte(content), &c)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return
	}
	fmt.Println(c)

}

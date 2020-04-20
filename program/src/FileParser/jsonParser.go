// Copyright (c) 2016-2020, Regents of the University of Arizona.
// Author:Wenkai Zheng<wenkaizheng@email.arizona.edu>
//        JiaCheng Yang <jiachengyang@email.arizona.edu>
// Read the network configuration file from json format
// Save data as a struct
// Print all data out in one line
package FileParser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readJson(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(f)
}

func GetJasonConfig(configPATH string, c interface{}) error {
	content, err := readJson(configPATH)
	if err != nil {
		fmt.Println("Cannot open " + configPATH + " : " + err.Error())
		return err
	}
	err = json.Unmarshal([]byte(content), c)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		return err
	}
	return nil

}

package Authentication

import (
	"Core"
	"FileParser"
	"Logging"
	"crypto/sha512"
	"fmt"
)

type userPasswordMap struct {
	userPassword map[string]string
	loaded       bool
}

func (s *userPasswordMap) Add(args ...interface{}) {
	strings, ok := args[0].([]string)
	if !ok {
		Logging.ErrorLogger.Fatal("for unknown reason, fail to parse elements in csv")
	}
	if len(strings) != 2 {
		Logging.ErrorLogger.Fatal("csv file contains more than 2 elements")
	}
	s.userPassword[strings[0]] = strings[1]
}

var record = userPasswordMap{make(map[string]string), false}

func LoadCSV(fileName string) {
	Logging.NormalLogger.Println("going to load CSV")
	if !record.loaded {
		FileParser.GetCSV(fileName, &record)
		record.loaded = true
	}
}

func Verify(username, password string) (bool, error) {
	Logging.NormalLogger.Println("going to verify given username and password")
	if !record.loaded {
		return false, nil //todo make a error
	}
	value, ok := record.userPassword[username]
	return ok && password == value, nil
}

func addSalt(s, salt1, salt2 string) string {
	return salt1 + s + salt2
}

func convert2Hex(arr []byte) string {
	return fmt.Sprintf("%X", arr)
}

func EncodePassword(password string) string {
	Logging.NormalLogger.Println("going to encode password")
	password = addSalt(password, "dlrC", "Ofsc")
	sha := sha512.New()
	sha.Write(Core.ConvertStringTOByte(password))
	return convert2Hex(sha.Sum(nil))
}

func EncodeUsername(username string) string {
	Logging.NormalLogger.Println("going to encode username")
	username = addSalt(username, "Bdho", "X643")
	sha := sha512.New()
	sha.Write(Core.ConvertStringTOByte(username))
	return convert2Hex(sha.Sum(nil))
}

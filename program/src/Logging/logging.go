package Logging

import (
	"log"
	"os"
)

var NormalLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
var ErrorLogger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)

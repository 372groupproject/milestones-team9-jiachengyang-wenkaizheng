package FileParser

import (
	"Logging"
	"encoding/csv"
	"io"
	"os"
)

type addable interface {
	Add(...interface{})
}

func GetCSV(fileName string, record addable) {
	f, err := os.Open(fileName)
	if err != nil {
		Logging.NormalLogger.Println("cannot open " + fileName)
		Logging.ErrorLogger.Fatal(err)
	}

	reader := csv.NewReader(f)
	for {
		result, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			Logging.NormalLogger.Println("encounter error while reading " + fileName)
			Logging.NormalLogger.Println()
			Logging.ErrorLogger.Fatal(err)
		}

		record.Add(result)
	}
}

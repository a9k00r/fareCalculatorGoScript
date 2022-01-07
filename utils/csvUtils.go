package utils

import (
	constant "Beat/constants"
	"encoding/csv"
	"log"
	"os"
)

func GetCsvReader() *csv.Reader {
	file, err := os.Open(constant.InputFilePath)
	if err != nil {
		log.Fatal(err)
	}
	return csv.NewReader(file)
}
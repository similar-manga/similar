package external

import (
	"encoding/csv"
	"log"
	"os"
)

func OpenCSVFileStream(path string) *os.File {
	file, err := os.Create(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func WriteToCSV(writer *csv.Writer, data []string) {
	err := writer.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}


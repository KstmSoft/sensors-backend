package services

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCSV(address string, body []string) {
	file, err := os.OpenFile(address+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed creating csv file: %s", err)
	}
	writer := csv.NewWriter(file)
	writer.Write(body)
	writer.Flush()
}

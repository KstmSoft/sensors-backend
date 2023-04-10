package services

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/viper"
)

func WriteCSV(filename string, body []string) {
	logsPath := viper.GetString("samples_folder")
	file, err := os.OpenFile(logsPath+"/"+filename+".csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Failed creating csv file: %s", err)
	}
	writer := csv.NewWriter(file)
	writer.Write(body)
	writer.Flush()
}
func ReadLogs(filename string) string {
	logsPath := viper.GetString("samples_folder")
	body, err := ioutil.ReadFile(logsPath + "/" + filename + ".csv")
	if err != nil {
		log.Printf("unable to read file: %v", err)
	}
	return string(body)
}

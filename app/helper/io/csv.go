package io

import (
	"encoding/csv"
	"os"
	"unico/app"
)

func ReadCSV(file string) [][]string {
	data, err := os.Open(file)
	if err != nil {
		app.ErrorLogger.Println(err)
	}
	defer data.Close()
	records, err := csv.NewReader(data).ReadAll()
	if err != nil {
		app.ErrorLogger.Println(err)
	}
	return records[1:]
}

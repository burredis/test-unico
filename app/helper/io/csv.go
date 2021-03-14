package io

import (
	"encoding/csv"
	"os"
	"fmt"
)

func ReadCSV(file string) [][]string {
	data, err := os.Open(file)
	if err != nil {
			fmt.Println(err)
	}
	defer data.Close()
	records, err := csv.NewReader(data).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	return records[1:]
}
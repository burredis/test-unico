package io

import (
	"os"
	"log"
	"fmt"
)

func Writefile(file string, data string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
}

func Removefile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println(err)
	}
}
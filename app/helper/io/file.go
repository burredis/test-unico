package io

import (
	"os"
	"unico/app"
)

func Writefile(file string, data string) {
	f, err := os.Create(file)
	if err != nil {
		app.ErrorLogger.Println(err)
	}
	defer f.Close()
	_, err = f.WriteString(data)
	if err != nil {
		app.ErrorLogger.Println(err)
	}
}

func Removefile(file string) {
	err := os.Remove(file)
	if err != nil {
		app.ErrorLogger.Println(err)
	}
}

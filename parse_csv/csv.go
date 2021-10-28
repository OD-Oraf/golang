package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

type Person struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lasname"`
	Address *Address `json:address`
}

type Address struct {
	City string `json:"city"`
	State string `json:"state"`
}
func main(){
	csvFile, _ := os.Open("data.csv")
	reader := csv.NewReader(csvFile)
	var people []Person
	//Loop until end of file
	for {
		line, error := reader.Read()
		if error == io.EOF {
			 else if error != nil {
			 log.Fatal(error)
		}
	}
}

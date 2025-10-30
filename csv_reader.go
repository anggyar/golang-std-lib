package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "anggyar,muhamad,yahya\n" + "pratama,putra,bagus\n" + "hello,world,go"

	var reader *csv.Reader = csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}

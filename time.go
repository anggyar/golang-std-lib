package main

import (
	"fmt"
	"time"
)

func main() {
	var now = time.Now()
	fmt.Println(now)

	var utc = time.Date(2025, time.June, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	formatter := "2006-01-02 15:04:05"

	value := "2025-06-01 10:00:00"
	valueTime, err := time.Parse(formatter, value)

	if err != nil {
		fmt.Println("Error", err.Error())
	} else {

		fmt.Println(valueTime)
	}

}

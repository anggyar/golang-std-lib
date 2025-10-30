package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"anggyar", "muhamad", "yahya"})
	_ = writer.Write([]string{"pratama", "putra", "bagus"})
	_ = writer.Write([]string{"hello", "world", "go"})

	writer.Flush()
}

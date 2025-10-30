package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("seafood", "food"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))

	fmt.Println(strings.Contains("Anggyar Muhamad Yahya", "Muhamad"))
	fmt.Println(strings.Split("Anggyar Muhamad Yahya", " "))
	fmt.Println(strings.ToLower("ANGGYAR MUHAMAD YAHYA"))
	fmt.Println(strings.Trim("    Anggyar Muhamad Yahya     ", " "))
	fmt.Println(strings.ReplaceAll("Anggyar Muhamad Yahya", "a", "o"))
}

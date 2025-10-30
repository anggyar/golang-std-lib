package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	value := "Anggyar Pratama"

	encode := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encode)

	decoded, err := base64.StdEncoding.DecodeString(encode)
	fmt.Println("hasil decoded", decoded)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}

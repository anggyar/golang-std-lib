package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	} else if id != "anggyar" {
		return NotFoundError
	}

	// Sukses
	return nil
}

func main() {

	err := GetById("anggyar")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Validation Error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Not found error")
		} else {
			fmt.Println("Unknown error:", err)
		}
	} else {
		fmt.Println("Success")
	}

}

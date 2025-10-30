package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {

	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value" + strconv.Itoa(i)
		data = data.Next()
	}

	// data := ring.New(5)
	// data.Value = "Anggyar"

	// data = data.Next()
	// data.Value = "Muhamad"

	// data = data.Next()
	// data.Value = "Yahya"

	// data = data.Next()
	// data.Value = "Developer"

	// data = data.Next()
	// data.Value = "Golang Enthusiast"

	data.Do(func(value any) {
		fmt.Println(value)
	})
}

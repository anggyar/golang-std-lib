package main

import (
	"fmt"
	"slices"
)

func main() {
	names := []string{"John", "Paul", "George"}
	values := []int{10, 20, 30}

	fmt.Println(slices.Max(values))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Contains(names, "Paul"))
	fmt.Println(slices.Index(names, "George"))
}

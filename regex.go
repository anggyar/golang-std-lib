package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex = regexp.MustCompile("b([a-z]+)i")

	fmt.Println(regex.MatchString("budi"))
	fmt.Println(regex.MatchString("b123i"))
	fmt.Println(regex.MatchString("babi"))
	fmt.Println(regex.MatchString("baci"))

	fmt.Println(regex.FindAllString("budi babi banci bini b4bu", 3))
}

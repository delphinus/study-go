package main

import "fmt"

func main() {
	var (
		a int
		b string
	)

	if a != 0 {
		fmt.Println("A!")
	} else {
		fmt.Println("not A!")
	}

	if b != "" {
		fmt.Println("B!")
	} else {
		fmt.Println("not B!")
	}
}

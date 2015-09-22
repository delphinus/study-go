package main

import "fmt"

func main() {
	i := new(int)

	*i = 123

	fmt.Println("*i:", *i)
}

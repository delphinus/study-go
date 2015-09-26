package main

import "fmt"

type Person struct { // 構造体型の宣言
	name string
	age  int
}

type Employee struct {
	id int
	Person
}

func main() {
	e := Employee{1, Person{"Jack", 28}}
	fmt.Println(e)
}

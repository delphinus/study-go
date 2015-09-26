package main

import "fmt"

type Person struct { // 構造体の宣言
	name string
	age  int
}

func main() {
	var p1 Person // 構造体リテラルを使用せず、フィールドを個別に初期化
	p1.name = "John"
	p1.age = 23

	p2 := Person{age: 31, name: "Tom"} // 構造体リテラルで初期化
	p3 := Person{"Jane", 42}           // 構造体リテラルで初期化
	p4 := &Person{"Mike", 36}          // ポインタも構造体リテラルで初期化可能
	fmt.Println(p1, p2, p3, p4)
}

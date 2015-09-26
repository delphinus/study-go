package main

import "fmt"

// int型からmyType型を宣言
type myType int

// 値を加算するメソッド
func (value *myType) add(increment myType) myType {
	*value += increment
	return *value
}

func main() {
	var i myType

	// 変数iに対し、addメソッドを呼び出し、3を加算
	fmt.Println(i.add(3))

	// メソッド値を取得。このメソッド値の型は「func(myType) myType」
	mv := i.add

	// メソッド値に対し、メソッドを呼び出し、再度3を加算
	fmt.Println(mv(3))
}

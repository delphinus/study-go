package main

import "fmt"

type embedded struct { // 埋め込まれる側の構造体
	i int
}

func (x embedded) doSomething() { // embedded型のメソッド
	fmt.Println("test.doSomething()")
}

type test struct { // 埋め込み先の構造体
	embedded // embedded型の埋め込み
}

func main() {
	var x test
	x.doSomething()  // embedded型に実装されているメソッドに、test型の値でアクセス
	fmt.Println(x.i) // embedded型のフィールドに、test型の値でアクセス
}

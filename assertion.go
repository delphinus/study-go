package main

import "fmt"

func main() {
	var i interface{} = "test" // 空インタフェースに、string型の値を格納
	var s string = i.(string)  // 型アサーションを使いstring型へ
	fmt.Println(s)             // 出力
}

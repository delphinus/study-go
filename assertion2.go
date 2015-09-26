package main

import "fmt"

func main() {
	var i interface{} = "test" // 空インタフェースに、string型の値を格納
	s1, ok := i.(string)       // 型アサーションに成功する例
	fmt.Println(s1, ok)
	s2, ok := i.(interface { // 型アサーションに失敗する例
		dummy()
	})
	fmt.Println(s2, ok)
}

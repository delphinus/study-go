package main

import "fmt"

func main() {
	s := []string{"a", "b", "c"} // スライスを作成
	test(s...)                   // スライスをそのまま渡すには...を付ける
	test("a", "b", "c")          // こうしたときと渡される値は一緒
}

// 可変長パラメータを持つ関数
func test(s ...string) {
	fmt.Println(len(s), s) // スライスの長さと値を出力
}

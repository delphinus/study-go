package main

import "fmt"

// 四則演算を行う関数の宣言
func calc(a int, b int) (add int, sub int, mul int, div float32) {
	// 戻り値を返す（足し算、引き算、掛け算、割り算）
	add = a + b
	sub = a - b
	mul = a * b
	div = float32(a) / float32(b)

	// この場合もreturn文は必要
	return
}

// main関数はパラメータ、戻り値ともに持たない
func main() {
	// calc関数を呼び出す
	add, sub, mul, div := calc(1, 2)

	// 関数から返された戻り値を出力
	fmt.Println(add, sub, mul, div)
}

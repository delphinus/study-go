package main

import "fmt"

type Calculator interface { // 演算インタフェース型
	Calculate(a int, b int) int // 関数の定義
}

type Add struct{} // 足し算型。フィールドは持たない
// Add型にCalculatorインタフェースのCalculate関数を実装
func (x Add) Calculate(a int, b int) int {
	return a + b // 足し算
}

type Sub struct{} // 引き算型。フィールドは持たない
// Sub型にCalculatorインタフェースのCalculate関数を実装
func (x Sub) Calculate(a int, b int) int {
	return a - b // 引き算
}

func main() {
	// Calculatorインタフェースを実装した型の変数を宣言
	var add Add
	var sub Sub
	var cal Calculator
	cal = add
	fmt.Println("和:", cal.Calculate(1, 2))
	cal = sub
	fmt.Println("差:", cal.Calculate(1, 2))
}

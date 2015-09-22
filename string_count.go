package main

import "fmt"
import "unicode/utf8" // 「utf8」パッケージをインポート

func main() {
  // string型の変数を宣言
  var ja string = "GO言語"

  // 文字数を出力
  fmt.Println(ja, "len:", utf8.RuneCountInString(ja))

  const G float32 = 9.80665
  const π float64 = 3.14159265
  const π2 float64 = π * 2

  fmt.Println("G:", G)
  fmt.Println("π :", π)
  fmt.Println("2 * π :", π2)
}

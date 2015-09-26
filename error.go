package main

import "fmt"

// マイエラー型
type MyError struct {
	message string // エラー詳細
}

func (err MyError) Error() string { // Error メソッドの実装
	return err.message // エラーメッセージを返す
}

func hex2int(hex string) (val int, err error) { // 16進数文字列をint型に変換
	for _, r := range hex { // 1文字ずつ取り出す
		val *= 16
		switch {
		case '0' <= r && r <= '9':
			val += int(r - '0')
		case 'a' <= r && r <= 'f':
			val += int(r-'a') + 10
		default:
			// 構造体リテラルを使ってMyErrorの値を作成し、エラー情報として返す
			return 0, MyError{"不正な文字です。: " + string(r)}
		}
	}
	return // 戻り値errには初期値であるnilが返る
}

func main() {
	fmt.Println(hex2int("1"))
	fmt.Println(hex2int("abcd"))
	fmt.Println(hex2int("z"))
}

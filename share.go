package main

import (
	"fmt"
	"os"
)

const goroutines = 10 // ゴルーチンの数

func main() {
	counter := make(chan int, 1) // 値の共用に使用するチャネルの作成
	for i := 0; i < goroutines; i++ {
		go func(counter chan int) { // ゴルーチン起動
			value := <-counter             // チャネルから値を取り出す
			value++                        // 値を加算
			fmt.Println("counter:", value) // 出力
			if value == goroutines {       // 全てのゴルーチンの処理が完了したら終了
				os.Exit(0)
			}
			counter <- value // 更新した値をチャネルに戻す
		}(counter)
	}
	counter <- 0 // チャネルに初期値を送信
	// 無限ループ
	// 環境変数「GOMAXPROCS」に2以上を指定する必要あり
	for {
	}
}

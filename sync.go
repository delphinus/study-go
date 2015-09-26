package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutines = 3 // 同時に行う処理数

func main() {
	c := make(chan int) // チャネル作成
	for i := 0; i < goroutines; i++ {
		go func(s chan<- int) { // 送信専用
			// ダミーの処理として乱数により0〜10秒待機
			time.Sleep(time.Duration(rand.Int63n(10)) * time.Second)
			fmt.Println("処理完了")
			// 処理完了をチャネルで伝える
			// 送信する値は何でも良い
			s <- 0
		}(c)
	}

	for i := 0; i < goroutines; i++ { // ゴルーチンの処理完了待ち
		<-c
	}
	fmt.Println("全て完了") // 完了
}

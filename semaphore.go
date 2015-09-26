package main

import (
	"fmt"
	"math/rand"
	"time"
)

const goroutines = 10  // ゴルーチンの数
const maxProcesses = 3 // 同時実行できる処理数上限

func main() {
	// セマフォ代わりのチャネルの作成（キャパシティに処理上限を設定）
	semaphore := make(chan int, maxProcesses)
	notify := make(chan int, goroutines) // 完了通知用チャネルの作成
	for i := 0; i < goroutines; i++ {
		go func(no int, semaphore chan int, notify chan<- int) {
			semaphore <- 0 // 自分の番が来るのを待つ
			// ダミーの処理として乱数により0〜3秒待機
			time.Sleep(time.Duration(rand.Int63n(3)) * time.Second)
			fmt.Println("処理完了", no)
			<-semaphore  // 待機中のゴルーチンに処理を明け渡す
			notify <- no // 処理完了を通知
		}(i, semaphore, notify)
	}
	for i := 0; i < goroutines; i++ { // ゴルーチンの処理完了待ち
		<-notify
	}
	fmt.Println("全て完了") // 完了
}

package main

import "fmt"

func main() {
	c := make(chan int) // チャネルの作成
	go func(s chan<- int) {
		for i := 0; i < 5; i++ { // チャネルへ0〜4の値を順番に送信
			s <- i
		}
		close(s) // チャネルのクローズ
	}(c)

	for { // 受信ループ
		value, ok := <-c // チャネルからの受信
		if !ok {         // チャネルがクローズされるとokにfalseが返される
			break
		}
		fmt.Println(value) // 受信した値を出力
	}
}

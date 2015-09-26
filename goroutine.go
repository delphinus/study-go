package main

import "fmt"
import "time"

func main() {
	fmt.Println("main:開始")
	fmt.Println("testを通常の関数として起動")
	test()
	fmt.Println("testをゴルーチンとして起動")
	go test()
	time.Sleep(3 * time.Second) // 3秒待つ
	fmt.Println("main:終了")
}

func test() { // ゴルーチンとして起動する関数
	for i := 0; i < 5; i++ {
		fmt.Println(i)              // 連番を出力
		time.Sleep(1 * time.Second) // 1秒待つ
	}
}

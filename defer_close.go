package main

import "fmt"
import "os"

func main() {
	// ファイルのオープン
	file, err := os.OpenFile("/var/log/system.log", os.O_WRONLY|os.O_CREATE, 0666)
	// オープンに失敗したときは終了
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// ファイルのクローズを遅延指定
	defer file.Close()
	// ファイルへテキスト出力
	file.WriteString("あいうえお")
}

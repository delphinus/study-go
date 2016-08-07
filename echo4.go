package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "改行を省きます。")
var sep = flag.String("s", " ", "引数の間に挟む文字列。")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

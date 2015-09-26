// ファイル先頭のコメントはドキュメントに出力されません

// sampleという名前のパッケージ
package sample

import "fmt"

// XXXは、値が「1234 * 2」の定数です。
const XXX = 1234 * 2

// YYYは、int型の変数です。
var YYY int

// ZZZはstringから作成した型です。
type ZZZ string

// test関数は何もしない関数です。
// エクスポートされていないのでドキュメントには出力されません。
func test(i) {
}

// Test関数は何もしない関数です。
func Test(x int, s string) {
}

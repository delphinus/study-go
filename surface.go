package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

const (
	width, height = 600, 300            // キャンパスの大きさ（画素数）
	cells         = 100                 // 格子の枡目の数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x 単位、及び、y 単位当たりの画素数
	zscale        = height * 0.4        // z 単位当たりの画素数
	angle         = math.Pi / 6         // x, y 軸の角度（=30 度）
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30 度), cos(30 度)

func main() {
	outFile, err := os.Create("surface.svg")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	fmt.Fprintf(outFile, strings.Join([]string{
		"<svg xmlns='http://www.w3.org/2000/svg'",
		"style='stroke: grey; fill: white; stroke-width: 0.7'",
		"width='%d' height='%d'>",
	}, " "), width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(outFile, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(outFile, "</svg>")
}

func corner(i, j int) (sx float64, sy float64) {
	// 枡目 (i, j) の角の点 (x, y) を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さ z を計算する。
	z := f(x, y)

	// (x, y, z) を 2-D SVG キャンバス (sx, sy) へ等角的に投影。
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	return
}

func f(x, y float64) (z float64) {
	r := math.Hypot(x, y) // (0, 0) からの距離
	z = math.Sin(r) / r
	return
}

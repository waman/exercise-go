// 【練習問題 3.3】
// 高さに基づいて個々のポリゴンに色付けし、頂点が赤 (#ff0000)
// となり谷が青 (#0000ff) になるようにしなさい。
package main

import (
	"fmt"
	"io"
	"log"
	"os"
	. "math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = Pi / 6
)

var sin30, cos30 = Sin(angle), Cos(angle)

func main() {
	var w io.Writer

	if len(os.Args) <= 1 {
		w = os.Stdout
	} else {
		// 引数があればファイルに出力（os パッケージのドキュメント参照）
		file, err := os.OpenFile(os.Args[1], os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0755)
		if err != nil {
			log.Fatal(err)
		}

		defer func() {
			if cErr := file.Close(); err == nil && cErr != nil {
				log.Fatal(err)
			}
		}()

		w = file
	}

	surface(w)
}

func surface(w io.Writer) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			top := Max(Max(Max(az, bz), cz), dz)
			var c string
			if top >= 0.0 {
				// top == 1 => #ff0000 : 赤
				// top == 0 => #ffffff : 白
				// Cbrt() はコントラストをつけるために入れてます。
				c = fmt.Sprintf("#ff%02x%02[1]x", int(255.0*(1-Cbrt(top))))
			} else {
				// top ==  0 => #ffffff : 白
				// top == -1 => #0000ff : 青
				c = fmt.Sprintf("#%02x%02[1]xff", int(255.0*(1+Cbrt(top))))
			}

			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c)
		}
	}
	fmt.Fprint(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := Hypot(x, y)
	return Sin(r) / r
}

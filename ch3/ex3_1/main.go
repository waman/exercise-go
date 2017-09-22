// 【連取問題 3.1】
// 関数 f が有限で内 float64 値を返すならば、SVG ファイルは不正な
// <polygon> 要素を含むことになります（もっとも、多くの SVG レンダラ
// はそれをうまく処理しますが）。不正なポリゴンをスキップするように
// プログラムを修正しなさい。
package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

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
			ax, ay := corner(i+1, j)
			// 有限でない値を返し得るのは y 座標だけ。
			if IsNotFinite(ay) {
				continue
			}

			bx, by := corner(i, j)
			if IsNotFinite(by) {
				continue
			}

			cx, cy := corner(i, j+1)
			if IsNotFinite(cy) {
				continue
			}

			dx, dy := corner(i+1, j+1)
			if IsNotFinite(dy) {
				continue
			}

			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprint(w, "</svg>")
}

// 引数が有限でない（正負の無限か NaN）なら true を、そうでないならfalse を返します。
func IsNotFinite(x float64) bool {
	return math.IsInf(x, 0) || math.IsNaN(x)
}

// 関数 f が有限でない値を返した場合、corner の第2返り値が有限でなくなります。
// 第3返り値にエラーを返す方が Go のコードっぽいですが、まだエラーの作り方が
// 現時点で出てきてないので、有限でない値をそのまま返しています
// （本文の surface と同じコード）。
func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

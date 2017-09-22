// 【練習問題 3.4】
// 1.7節のリサジューの例での方法に従って、面を計算して SVG データを
// クライアントに書き出すウェブサーバを作成しなさい。サーバは次のように
// Content-Type を設定しなければなりません。
//
//   w.Header().Set("Content-Type", "image/svg+xml")
//
// （このステップはリサジューの例では必要ありませんでした。それは、サーバが
// 標準的なヒューリスティクスを使ってレスポンスの最初の512バイトから PNG
// などの共通形式を認識し、適切なヘッダーを生成しているからです。）HTTP
// リクエストのパラメータとして、高さ、幅、色などの値をクライアントが指定
// できるようにしなさい。
package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	. "math"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("SVG Server starts...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	width, height, cells, xyrange, xyscale, zscale, angle, c := getParameters(r)

	fwidth, fheight, fcells := float64(width), float64(height), float64(cells)
	sinPhi, cosPhi := Sincos(angle) // math.Sincos() は sin と cos をまとめて計算する

	corner := func(i, j int) (float64, float64) {
		x := xyrange * (float64(i)/fcells - 0.5)
		y := xyrange * (float64(j)/fcells - 0.5)

		z := f(x, y)

		sx := fwidth/2 + (x-y)*cosPhi*xyscale
		sy := fheight/2 + (x+y)*sinPhi*xyscale - z*zscale
		return sx, sy
	}

	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' stroke='%s'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, c)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

const (
	defaultWidth, defaultHeight = 600, 320
	defaultCells                = 100
	defaultXYRange              = 30.0
	defaultAngle                = Pi / 6
	defaultColor                = "#000000"
)

// 返り値: width, height, cells, xyrange, xyscale, zscale, angle, color
// マップも構造体も使い方が出てきてないので返り値を列挙してます ^^;)
func getParameters(r *http.Request) (int, int, int, float64, float64, float64, float64, string) {

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	// width
	width, err := strconv.Atoi(r.Form.Get("width"))
	if err != nil {
		width = defaultWidth
	}

	// height
	height, err := strconv.Atoi(r.Form.Get("height"))
	if err != nil {
		height = defaultHeight
	}

	// cells
	cells, err := strconv.Atoi(r.Form.Get("cells"))
	if err != nil || cells > 100 {
		cells = defaultCells
	}

	// xyrange  strconv.ParseFloat() は整数に対する strconv.AtoI と同じような関数
	xyrange, err := strconv.ParseFloat(r.Form.Get("xyrange"), 64)
	if err != nil {
		xyrange = defaultXYRange
	}

	// xyscale
	xyscale, err := strconv.ParseFloat(r.Form.Get("xyscale"), 64)
	if err != nil {
		xyscale = float64(width) / 2 / xyrange
	}

	// zscale
	zscale, err := strconv.ParseFloat(r.Form.Get("zscale"), 64)
	if err != nil {
		zscale = float64(height) * 0.4
	}

	// angle
	angle, err := strconv.ParseFloat(r.Form.Get("angle"), 64)
	if err != nil {
		angle = defaultAngle
	}

	// color
	c := r.Form.Get("color")
	if len(c) == 0 {
		c = defaultColor
	} else if _, err := strconv.ParseInt(c, 16, 64); err == nil {
		c = "#" + c
	}

	return width, height, cells, xyrange, xyscale, zscale, angle, c
}

func f(x, y float64) float64 {
	r := Hypot(x, y)
	return Sin(r) / r
}

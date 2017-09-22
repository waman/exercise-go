// 【練習問題 3.9】
// フラクタルをレンダリングして画像データをクライアントへ書き出す
// ウェブサーバを作成しなさい。HTTP リクエストへのパラメータとして、
// クライアントが x, y, 倍率値を指定できるようにしなさい。
package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Mandelbrot Server starts...")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// クエリとしては中心の座標 x, y と倍率 scale を許容しています。
// 描画される領域は (x-scale, y-scale), (x+scale, y+scale)
func handler(w http.ResponseWriter, r *http.Request) {

	const res = 1024 // == width, height

	cx, cy, scale := getParameters(r)

	xmin, ymin := cx-scale, cy-scale
	del := scale * 2 / res

	img := image.NewRGBA(image.Rect(0, 0, res, res))
	for py := 0; py < res; py++ {
		y := float64(py)*del + ymin
		for px := 0; px < res; px++ {
			x := float64(px)*del + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(w, img)
}

func getParameters(r *http.Request) (float64, float64, float64) {

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}

	// x 中心の x 座標
	x, err := strconv.ParseFloat(r.Form.Get("x"), 64)
	if err != nil {
		x = 0
	}

	// y 中心の y 座標
	y, err := strconv.ParseFloat(r.Form.Get("y"), 64)
	if err != nil {
		y = 0
	}

	// scale スケール
	scale, err := strconv.ParseFloat(r.Form.Get("scale"), 64)
	if err != nil {
		scale = 2
	}

	return x, y, scale
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}

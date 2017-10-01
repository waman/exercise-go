// 【練習問題 3.7】
// 別の単純なフラクタルは z^4 - 1 = 0 などの方程式に対する複素数解を
// 求めるためにニュートン法を使います。四つの根の一つに近づくのに必要な
// 繰返し回数で各開始点にグラデーションを付けなさい。それが近づいている
// 根ごとに各点に色付けしなさい。
package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"math/cmplx"
	"os"
)

func main() {

	var w io.Writer

	if len(os.Args) <= 1 {
		w = os.Stdout
	} else {
		// 引数があればファイルに出力（os パッケージのドキュメント参照）
		file, err := os.Create(os.Args[1])
		if err != nil {	log.Fatal(err) }

		defer func() {
			if err := file.Close(); err != nil { log.Fatal(err) }
		}()

		w = file
	}

	outputImage(w)
}

func outputImage(w io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, julia(z))
		}
	}
	png.Encode(w, img)
}

// y = f(x) の x = x_1 における接線の方程式は
//
//   y - f(x_1) = f'(x_1)(x - x_1)
//
// なので、この接線上の点で y = 0 となる x の値は
//
//   x = x_1 - f(x_1)/f'(x_1)
//
// よってニュートン法による複素数平面上の点の移動は、以下の漸化式
//
//   z_{n+1} = z_1 - f(z_n)/f'(z_n)
//
// で与えられます。特に f(z) = z^4 - 1 のとき
//
//   z_{n+1} = 3z_n/4 + 1/4z_n^3
//
// となります。
//
// この方法で得られる図形はジュリア集合と呼ばれるようなので（正確な定義は他所参照）
// メソッド名は julia としています。
func julia(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	for n := uint8(0); n < iterations; n++ {

		if nearlyEqual(z, 1) {
			return color.Gray{Y: 255 - contrast*n}

		} else if nearlyEqual(z, 1i) {
			return color.RGBA{R: 255 - contrast*n, G: 0xff, B: 0xff, A: 0xff}

		} else if nearlyEqual(z, -1) {
			return color.RGBA{R: 0xff, G: 255 - contrast*n, B: 0xff, A: 0xff}

		} else if nearlyEqual(z, -1i) {
			return color.RGBA{R: 0xff, G: 0xff, B: 255 - contrast*n, A: 0xff}
		}

		z = z*3/4 + 1.0/(z*z*z*4)
	}
	return color.Black
}

// z と方程式 z^4 - 1 = 0 の解 (z = 1, i, -1, -i) との距離が
// delta 未満になったら繰り返しをやめる。
const delta = 0.1

func nearlyEqual(z0, z1 complex128) bool {
	return cmplx.Abs(z0-z1) < delta
}

package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"log"
	"os"
	. "math/big"
)

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

	outputImageFloat(w)
}

func outputImageFloat(w io.Writer) {
	minusHalf := newFloat().Quo(fl(-1), fl(2))

	xmin, xmax := minusHalf, fl(0)
	ymin, ymax := fl(-1), minusHalf
	width, height := 512, 512

	dx, dy := newFloat(), newFloat()
	dx.Quo(dx.Sub(xmax, xmin), fl(width))  // dx = (xmax-xmin)/width
	dy.Quo(dy.Sub(ymax, ymin), fl(height)) // dy = (ymax-ymin)/height

	setInt := func(x *Float, i int) *Float {
		return x.SetInt64(int64(i))
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	x, y := newFloat(), newFloat()
	for py := 0; py < height; py++ {
		y.Add(y.Mul(setInt(y, py), dy), ymin) // y = py*dy + ymin
		for px := 0; px < width; px++ {
			x.Add(x.Mul(setInt(x, px), dx), xmin) // x = px*dx + xmin
			img.Set(px, py, mandelbrotFloat(x, y))
		}
	}
	png.Encode(w, img)
}

func newFloat() *Float {
	return new(Float).SetPrec(128)
}

func fl(i int) *Float {
	return newFloat().SetInt64(int64(i))
}

func mandelbrotFloat(x, y *Float) color.Color {
	const iterations = 200
	const contrast = 15
	two, four := fl(2), fl(4) // = 2^2 半径の自乗

	// mainRat.go で同名の関数を定義するので、
	// 名前の衝突を避けるためにローカル関数にしています。

	// マンデルブロ集合を作る漸化式 z = z*z + c
	// 実部・虚部を分けて書けば z = x + yi, c = a + bi として
	//   newX = x^2 - y^2 + a
	//   newY = 2xy + b
	f := func(x, y, a, b *Float) (*Float, *Float) {
		u, v, y2 := newFloat(), newFloat(), newFloat()
		u.Add(u.Sub(u.Mul(x, x), y2.Mul(y, y)), a)
		v.Add(v.Mul(v.Mul(x, y), two), b)
		return u, v
	}

	// x + yi の絶対値の自乗を返す。
	complexAbs2 := func(x, y *Float) *Float {
		abs2 := newFloat()
		abs2.Add(abs2.Mul(x, x), newFloat().Mul(y, y))
		return abs2
	}

	u, v, dif := newFloat(), newFloat(), newFloat()
	for n := uint8(0); n < iterations; n++ {
		u, v = f(u, v, x, y)
		if dif.Sub(complexAbs2(u, v), four).Sign() > 0 {
			return color.Gray{Y: 255 - contrast*n}
		}
	}
	return color.Black
}

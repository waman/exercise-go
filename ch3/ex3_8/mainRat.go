package main

import (
	"image/png"
	"os"
	. "math/big"  // 接頭辞なしで書けるようにするため
	"image/color"
	"image"
)

func main(){
	xmin, xmax    := rat(-1, 2), rat(0, 1)
	ymin, ymax    := rat(-1, 1), rat(-1, 2)
	width, height := 512, 512

	dx, dy := new(Rat), new(Rat)
	dx.Quo(dx.Sub(xmax, xmin), rat(width, 1))   // dx = (xmax-xmin)/width
	dy.Quo(dy.Sub(ymax, ymin), rat(height, 1))  // dy = (ymax-ymin)/height

	setInt := func(x *Rat, i int) *Rat {
		return x.SetInt64(int64(i))
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	x, y := new(Rat), new(Rat)
	for py := 0; py < height; py++ {
		y.Add(y.Mul(setInt(y, py), dy), ymin)  // y = py*dy + ymin
		for px := 0; px < width; px++ {
			x.Add(x.Mul(setInt(x, px), dx), xmin)  // x = px*dx + xmin
			img.Set(px, py, mandelbrotRat(x, y))
		}
	}
	png.Encode(os.Stdout, img)
}

func rat(i, j int) *Rat{
	return NewRat(int64(i), int64(j))
}

func mandelbrotRat(x, y *Rat) color.Color {
	const iterations = 200
	const contrast   = 15
	two, four := rat(2, 1), rat(4, 1)  // = 2^2 半径の自乗

	// mainRat.go で同名の関数を定義するので、
	// 名前の衝突を避けるためにローカル関数にしています。

	// マンデルブロ集合を作る漸化式 z = z*z + c
	f := func(x, y, a, b *Rat) (*Rat, *Rat){
		u, v, y2 := new(Rat), new(Rat), new(Rat)
		u.Add(u.Sub(u.Mul(x, x), y2.Mul(y, y)), a)
		v.Add(v.Mul(v.Mul(x, y), two), b)
		return u, v
	}

	// x + yi の絶対値の自乗を返す。
	complexAbs2 := func(x, y *Rat) *Rat {
		abs2 := new(Rat)
		abs2.Add(abs2.Mul(x, x), new(Rat).Mul(y, y))
		return abs2
	}

	u, v, dif := new(Rat), new(Rat), new(Rat)
	for n := uint8(0); n < iterations; n++ {
		u, v = f(u, v, x, y)
		if dif.Sub(complexAbs2(u, v), four).Sign() > 0 {
			return color.Gray{ Y:255 - contrast*n }
		}
	}
	return color.Black
}

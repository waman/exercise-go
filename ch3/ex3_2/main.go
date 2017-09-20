// surface は 3-D 面の関数の SVG レンダリングを計算します。
package main

import (
	"math"
	"fmt"
	"os"
	"log"
	"io"
)

const(
	width, height = 600, 320
	cells   = 100
	xyrange = 30.0
	xyscale = width / 2 / xyrange
	zscale  = height * 0.001
	angle   = math.Pi/6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// 引数に 1, 2, 3 を指定して、標準出力に対応するグラフの SVG を書き出す。
//   1. 鶏卵の箱
//   2. モーグルのこぶ
//   3. 乗馬用の鞍
func main(){
	// 描画する関数
	var f func(float64, float64)float64

  if len(os.Args) == 1 {
		log.Fatal("引数に 1, 2, 3 のいずれかを指定してください。")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "1": f = eggCrate
	case "2": f = mogul
	case "3": f = saddle
	default:
		log.Fatalf("引数に 1, 2, 3 のいずれかを指定してください: %s", os.Args[1])
		os.Exit(1)
	}

	var w io.Writer
	if len(os.Args) == 2 {
		w = os.Stdout
	}else {
		// 2つ目の引数があればファイルに出力（os パッケージのドキュメント参照）
		file, err := os.OpenFile(os.Args[2], os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0755)
		if err != nil { log.Fatal(err) }

		defer func(){
			if cErr := file.Close(); err == nil && cErr != nil {
				log.Fatal(err)
			}}()

		w = file
	}

	surface(w, f)
}

func surface(w io.Writer, f func(float64, float64)float64){

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' " +
		"style='stroke: grey; fill: white; stroke-width:0.7' " +
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1,j+1, f)
			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprint(w, "</svg>")
}

func corner(i, j int, f func(float64, float64)float64)(float64, float64){
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

// 鶏卵の箱。
// Egg Create Function x^2 + y^2 + 25(sin^2 x + sin^2 y) を
// 図を見易くするために少し改良。
func eggCrate(x, y float64) float64 {
	sx, sy := math.Sin(x), math.Sin(y)
	return x*x + y*y + 50*(sx*sx + sy*sy)
}

// モーグルのこぶ。
// sin x sin y の適当なスケール変換。
// Egg Create Function の第3項と（定数倍と定数項を除いて）実質的に同じ。
func mogul(x, y float64) float64 {
	return 80*math.Sin(x/2)*math.Sin(y/2)
}

// 鞍点の図によく用いられる関数 x^2 - y^2 を
// 図を見やすくするために少し改良。
func saddle(x, y float64) float64 {
	return x*x - 3*y*y
}

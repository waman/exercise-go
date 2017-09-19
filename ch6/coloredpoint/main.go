// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	. "github.com/waman/exercise-go/ch6/geometry"  // 接頭辞なしで import
	"image/color"
	"fmt"
)

type ColoredPoint struct {
	Point
	Color color.RGBA
}

// Point へのポインタを埋め込むバージョン
type ColoredPoint2 struct {
	*Point
	Color color.RGBA
}

func main(){
	// ColoredPoint
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)  // "1"

	cp.Point.Y = 2
	fmt.Println(cp.Y)        // "2"

	red  := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	var q = ColoredPoint{Point{5, 4}, blue}

	// p.Distance(q)  // コンパイルエラー
	fmt.Println(p.Distance(q.Point))  // "5"

	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))  // "10"

	// ColoredPoint2
	r := ColoredPoint2{&Point{1, 1}, red}
	s := ColoredPoint2{&Point{5, 4}, blue}

	fmt.Println(r.Distance(*s.Point))  // "5"

	s.Point = r.Point
	r.ScaleBy(2)
	fmt.Println(*r.Point, *s.Point)  // "{2, 2} {2, 2}"
}
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	. "github.com/waman/exercise-go/ch6/geometry"
	"fmt"
	)

func main(){
	// メソッド値
	p := Point{1, 2}
	q := Point{4, 6}

	distanceFromP := p.Distance
	fmt.Println(distanceFromP(q)) // "5

	var origin Point                   // {0, 0}
	fmt.Println(distanceFromP(origin)) // ~ √5

	scaleP := p.ScaleBy
	scaleP(2)  // p == {2, 4}
	scaleP(3)  // p == {6, 12}
	scaleP(10) // p == {60, 120}

	// メソッド式
	r := Point{1, 2}
	s := Point{4, 6}

	distance := Point.Distance
	fmt.Printf("%T\n", distance) // "func(Point, Point) float64"
	fmt.Println(distance(r, s))  // "5"

	scale := (*Point).ScaleBy
	fmt.Printf("%T\n", scale)  // "func(*Point, float64)"
	scale(&r, 2)
	fmt.Println(r) // {2, 4}
}

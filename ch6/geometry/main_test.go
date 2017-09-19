// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package geometry

import (
	"testing"
)

func TestPoint(t *testing.T){
	p := Point{1, 2}
	q := Point{4, 6}
	if Distance(p, q) != 5 { t.Fail() }
	if p.Distance(q) != 5 { t.Fail()}
}

func TestScaleBy(t *testing.T){
	// 呼び出し その1
	p := &Point{1, 2}
	p.ScaleBy(2)
	if !(p.X == 2 && p.Y == 4) { t.Fail() }

	// 呼び出し その2 （普通やらない）
	q := Point{1, 2}
	pptr := &q
	pptr.ScaleBy(2)
	if !(q.X == 2 && q.Y == 4) { t.Fail() }

	// 呼び出し その3 （普通やらない）
	r := Point{1, 2}
	(&r).ScaleBy(2)
	if !(r.X == 2 && r.Y == 4) { t.Fail() }

	// 呼び出し その4
	s := Point{1, 2}
	s.ScaleBy(2)
	if !(s.X == 2 && s.Y == 4) { t.Fail() }
}

func TestPath(t *testing.T){
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	if perim.Distance() != 12 { t.Fail() }
}

func TestTranslate(t *testing.T){
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
	}

	v := Point{2, 3}

	perim.TranslateBy(v, true)

	if !(perim[0] == Point{3, 4} &&
		   perim[1] == Point{7, 4} &&
	     perim[2] == Point{7, 7}){ t.Fail() }
}
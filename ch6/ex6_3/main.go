// 【練習問題 6.3】
// (*IntSet).UnionWith はワード単位のビット和算子である | を使って
// 二つのセットの和集合を計算しています。セット操作に対応するメソッド
// IntersectWith, DifferenceWith, SymmetricDifference を実装
// しなさい。（二つの集合の対称差は、どちらかの集合にはあるが、両方には
// ない要素を含む集合です。）
package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

// IntersectWith は、s と t の積集合（共通部分）を s に設定します。
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = s.words[:i]
			return
		}
	}
	s.words = s.words[:len(t.words)]
}

// Difference は、s から t に含まれる要素を除いた集合を s に設定します。
func (s *IntSet) Difference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			break
		}
	}
}

// SymmetricDifference は、s と t の対称差を s に設定します。
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func main() {
	// UnionWith
	var x0, x1 IntSet
	x0.AddAll(1, 9, 43, 144, 1222, 1223)
	x1.AddAll(9, 42, 144, 1000)

	x0.IntersectWith(&x1)
	fmt.Println(x0.String()) // "{9 144}"

	var y0, y1 IntSet
	y1.AddAll(9, 42, 144, 1000)
	y0.AddAll(1, 9, 43, 144, 1222, 1223)

	y0.IntersectWith(&y1)
	fmt.Println(y0.String()) // "{9 144}"

	// Difference
	var z0, z1 IntSet
	z0.AddAll(1, 9, 43, 144, 1222, 1223)
	z1.AddAll(9, 42, 144, 1000)

	z0.Difference(&z1)
	fmt.Println(z0.String()) // "{1 43 1222 1223}"

	// SymmetricDifference
	var t0, t1 IntSet
	t0.AddAll(1, 9, 43, 144, 1222, 1223)
	t1.AddAll(9, 42, 144, 1000)

	t0.SymmetricDifference(&t1)
	fmt.Println(t0.String()) // "{1 42 43 1000 1222 1223}"
}

// ***** 以下は本文のコードより *****
// Add はセットに負ではない値 x を追加します。
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// String は "{1 2 3}" の形式の文字列としてセットを返します。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//***** 以下は練習問題 6.2 より *****
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

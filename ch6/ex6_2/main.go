// 【練習問題 6.2】
// s.AddAll(1, 2, 3) などのように値のリストが追加可能である可変引数
// (*IntSet).AddAll(...int) メソッドを定義しなさい。
package main

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
}

func main() {
	var x IntSet
	x.AddAll(1, 144, 9)
	fmt.Println(x.String()) // "{1, 9, 144}"

	var y IntSet
	y.Add(1)
	y.Add(144)
	y.Add(9)

	y.AddAll(2, 3, 10)
	fmt.Println(y.String()) // "{1, 2, 3, 9, 144}"
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

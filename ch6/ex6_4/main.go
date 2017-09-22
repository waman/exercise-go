// 【練習問題 6.4】
// range ループでの繰り返しに適した、セットの要素を含むスライスを返す
// Elems メソッドを追加しなさい。
package main

import (
	"bytes"
	"fmt"
	"github.com/waman/exercise-go/ch2/popcount"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		for j := 0; j < 64; j++ {
			if (1 << uint(j)) & word != 0 {
				elems = append(elems, i*64+int(j))
			}
		}
	}
	return elems
}

func main() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.Elems()) // "{1 9 144}"

	var y IntSet
	y.Add(9)
	y.Add(42)
	fmt.Println(y.Elems()) // "{9 42}"

	var z IntSet
	z.Add(0)
	z.Add(1)
	z.Add(13)
	z.Add(63)
	z.Add(64)
	z.Add(1223)
	z.Add(1224)
	fmt.Println(z.Elems()) // "{0 1 13 63 64 1223 1224}"
}


// ***** 以下は本文のコードより *****
func (s *IntSet) Len() int {
	var n = 0
	for _, word := range s.words {
		n += popcount.PopCount(word)
	}
	return n
}

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

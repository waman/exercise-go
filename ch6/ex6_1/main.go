// 【練習問題 6.1】
// これらの追加メソッドを実装しなさい。
//
//   func (*IntSet) Len() int      // 要素数を返します
//   func (*IntSet) Remove(x int)  // セットから x を取り除きます
//   func (*IntSet) Clear()        // セットからすべての要素を取り除きます
//   func (*IntSet) Copy() *IntSet // セットのコピーを返します。
package main

import (
	"bytes"
	"fmt"
	"github.com/waman/exercise-go/ch2/popcount"
)

type IntSet struct {
	words []uint64
}

func (s *IntSet) Len() int {
	var n = 0
	for _, word := range s.words {
		n += popcount.PopCount(word)
	}
	return n
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	s.words[word] &^= (1 << bit)

	for i := len(s.words) - 1; i >= 0; i-- {
		if s.words[i] != 0 {
			s.words = s.words[:i+1]
			return
		}
	}

	s.words = nil
}

func (s *IntSet) Clear() {
	s.words = nil
}

func (s *IntSet) Copy() *IntSet {
	var newWords = make([]uint64, len(s.words))
	copy(newWords, s.words)
	return &IntSet{newWords}
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)

	y.Add(9)
	y.Add(42)

	// Len()
	fmt.Println(x.Len()) // 3
	fmt.Println(y.Len()) // 2

	// Remove()
	x.Remove(1)
	fmt.Println(x.String()) // "{9, 144}"
	x.Remove(10)
	fmt.Println(x.String()) // "{9, 144}"

	// Clear()
	x.Clear()
	fmt.Println(x.String()) // "{}"

	// Copy()
	var z = y.Copy()
	fmt.Println(z.String()) // "{9, 42}"
	y.Clear()
	fmt.Println(y.String()) // "{}"
	fmt.Println(z.String()) // "{9, 42}"
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

// 【練習問題 6.5】
// IntSet で使われている個々のワードの型は uint64 ですが、64ビット演算は
// 32ビットプラットフォームでは非効率かもしれません。プラットフォームに
// 対して最も効率的な符号なし整数である uint 型を使うようにプログラムを
// 修正しなさい。64で割る代わりに、uint の実質的サイズのバイト数である
// 32あるいは64を保持する定数を定義しなさい。そのためには、おそらくかなり
// 賢い式 32 << (^uint(0) >> 63) を使えます。
package main

import (
	"bytes"
	"fmt"
)

const base int = 32 << (^uint(0) >> 63)

// 第2章の PopCount より
// pc[i] は i のポピュレーションカウントです。
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount は x のポピュレーションカウント（1が設定されているビット）を返します。
func popCount32(x uint) int {
	y := uint32(x)
	return int(pc[byte(y>>(0*8))] +
		pc[byte(y>>(1*8))] +
		pc[byte(y>>(2*8))] +
		pc[byte(y>>(3*8))])
}

func popCount64(x uint) int {
	y := uint64(x)
	return int(pc[byte(y>>(0*8))] +
		pc[byte(y>>(1*8))] +
		pc[byte(y>>(2*8))] +
		pc[byte(y>>(3*8))] +
		pc[byte(y>>(4*8))] +
		pc[byte(y>>(5*8))] +
		pc[byte(y>>(6*8))] +
		pc[byte(y>>(7*8))])
}

// IntSet は負ではない小さな整数のセットです。
// そのゼロ値は空セットを表しています。
type IntSet struct {
	words []uint
}

// Has は負ではない値 x をセットが含んでいるか否かを報告します。
func (s *IntSet) Has(x int) bool {
	word, bit := x/base, uint(x%base)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add はセットに負ではない値 x を追加します。
func (s *IntSet) Add(x int) {
	word, bit := x/base, uint(x%base)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith は、s と t の和集合を s に設定します。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String は "{1 2 3}" の形式の文字列としてセットを返します。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < base; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", base*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func (s *IntSet) Len() int {
	var n = 0
	var f func(uint)int
	if base == 64 {
		f = popCount64
	}else{
		f = popCount32
	}

	for _, word := range s.words {
		n += f(word)
	}
	return n
}

func (s *IntSet) Remove(x int) {
	word, bit := x/base, uint(x%base)
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
	var newWords = make([]uint, len(s.words))
	copy(newWords, s.words)
	return &IntSet{newWords}
}

func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
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

func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		for j := 0; j < base; j++ {
			if (1 << uint(j)) & word != 0 {
				elems = append(elems, i*base+int(j))
			}
		}
	}
	return elems
}

func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"

	fmt.Println(x.Len())
}

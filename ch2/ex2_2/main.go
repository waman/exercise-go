// 【練習問題 2.2】
// コマンドライン引数、もしくはコマンドライン引数が指定されなかった場合には
// 標準入力から数値を読み込む、cf に似た単位変換プログラムを書きなさい。
// 各数値は、温度なら華氏と摂氏で、長さならフィートとメートルで、重さなら
// ポンドとキログラムでといった具合に核種単位へ変換しなさい。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	MetrePerFeet     = 0.3048
	KilogramPerPound = 0.45359237
)

// Temperature
type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32)) * 5 / 9 }

// Length
type Metre float64
type Feet float64

func (d Metre) String() string { return fmt.Sprintf("%gm", d) }
func (f Feet) String() string  { return fmt.Sprintf("%gft", f) }

func MToFt(d Metre) Feet  { return Feet(d / MetrePerFeet) }
func FtToM(ft Feet) Metre { return Metre(ft * MetrePerFeet) }

// Mass
type Kilogram float64
type Pound float64

func (m Kilogram) String() string { return fmt.Sprintf("%gkg", m) }
func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }

func KgToLb(m Kilogram) Pound { return Pound(m / KilogramPerPound) }
func LbToKg(p Pound) Kilogram { return Kilogram(p * KilogramPerPound) }

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		input := bufio.NewScanner(os.Stdin)

		fmt.Print("> ")
		for input.Scan() {
			calculateConversions(input.Text())
			fmt.Print("> ")
		}
	} else {
		for _, arg := range args {
			calculateConversions(arg)
		}
	}
}

func calculateConversions(s string) {
	x, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	// Temperature
	c := Celsius(x)
	f := Fahrenheit(x)
	fmt.Printf("%s = %s, %s = %s\n", c, CToF(c), f, FToC(f))

	// Length
	d := Metre(x)
	ft := Feet(x)
	fmt.Printf("%s = %s, %s = %s\n", d, MToFt(d), ft, FtToM(ft))

	// Mass
	m := Kilogram(x)
	p := Pound(x)
	fmt.Printf("%s = %s, %s = %s\n", m, KgToLb(m), p, LbToKg(p))

	fmt.Println()
}

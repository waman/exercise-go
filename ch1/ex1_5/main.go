// 【練習問題 1.5】
// もっともらしくするために、リサジュープログラムのカラーパレットを背景を
// 黒として緑の線になるように修正しなさい。ウエブカラー #RRGGBB を作成する
// ために color.RGBA{0xRR, 0xGG, 0xBB, 0xff} を使いなさい。
// ここで16進表記のそれぞれの値は、画素の赤、緑、青の成分の強度を表して
// います。
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette2 = []color.Color{color.Black, color.RGBA{0x00, 0xff, 0x00, 0xff}}

const (
	backgroundIndex = 0
	lineIndex       = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	var w io.Writer

	if len(os.Args) <= 1 {
		w = os.Stdout
	} else {
		// 引数があればファイルに出力（os パッケージのドキュメント参照）
		file, err := os.Create(os.Args[1])
		if err != nil { log.Fatal(err) }

		defer func() {
			if err := file.Close(); err != nil { log.Fatal(err) }
		}()

		w = file
	}

	lissajous(w)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette2)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), lineIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

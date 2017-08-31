// 【練習問題 1.6】
// リサジュープログラムを修正して palette にもっと値を追加し、
// 何らかの興味深い方法で SetColorIndex の第3引数を変更して
// 複数の色で画像を生成するようにしなさい。
package main

import(
	"image"
	"math/rand"
	"time"
	"os"
	"io"
	"image/gif"
	"math"
	"image/color/palette"
)

func main(){
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func lissajous(out io.Writer){
	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		// パレットに色を追加するのが面倒なので
		// 標準 API の image/color/palette にある Plan9 を使用
		img := image.NewPaletted(rect, palette.Plan9)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			// 時間ごと (i ごと) に色を変化
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(i))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, & anim)
}

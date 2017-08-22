// 【練習問題 1.12】
// リサージュ図形のサーバを修正して、URL からパラメータ値を読み取るように
// しなさい。たとえば、http://localhost:8000/?cycles=20 のような URL
// では、周回の回数をデフォルトの5ではなく20に設定するようにしなさい。
// 文字列パラメータを整数へ変換するために strconv.Atoi 関数を使いなさい。
// その変換関数のドキュメントは go doc strconv.Atoi で見ることができます。
package main

import (
	"net/http"
	"log"
	"image/gif"
	"image"
	"math"
	"math/rand"
	"image/color"
	"strconv"
)

var palette = []color.Color{color.White, color.Black}

const(
	whiteIndex = 0
	blackIndex = 1
)

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	const (
		res = 0.001
		size = 100
		nframes = 64
		delay = 8
	)

	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles, err := strconv.Atoi(r.Form.Get("cycles"))
	if err != nil {
		cycles = 5  // クエリで cycles が指定されてなければ5とする
	}

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		// cycles を float64 にキャストする必要があるもよう
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(w, & anim)  // ResponseWriter である w へ書き出す
}

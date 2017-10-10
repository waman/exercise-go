// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// sync.Once で遅延初期化
package main

import (
	"image"
	"sync"
)

var loadIconsOnce sync.Once
var icons map[string]image.Image

func loadIcons(){
	icons = map[string]image.Image{
		"spades.png"  : loadIcon("spades.png"),
		"hearts.png"  : loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png"   : loadIcon("clubs.png"),
	}
}

func loadIcon(name string) image.Image {
	panic("implement me.")  // アイコンの画像を読み込んで返す
}


// 並行的に安全
func Icon(name string) image.Image {
	loadIconsOnce.Do(loadIcons)
	return icons[name]
}

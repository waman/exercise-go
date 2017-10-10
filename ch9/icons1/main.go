// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import "image"

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


// 注意：並行的に安全でない！
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"fmt"
	"flag"
)

// du1 コマンドは、ディレクトリ内のファイルのディスク容量を計算します。
func main(){
	// 最初のディレクトリを決める
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// ファイルツリーを走査する
	fileSizes := make(chan int64)
	go func(){
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// 結果を表示する
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}

// walkDir は dir をルートとするファイルツリーをたどり、
// 見つかったファイルのそれぞれの大きさを fileSizes に送ります。
func walkDir(dir string, fileSizes chan<- int64){
	for _, entry := range dirents(dir){
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}

// 表示を GB から MB に変えています。
func printDiskUsage(nfiles, nbytes int64){
	fmt.Printf("%d files  %1.f MB\n", nfiles, float64(nbytes)/1e6)
}
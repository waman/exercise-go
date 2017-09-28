// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"path/filepath"
	"os"
	"io/ioutil"
	"fmt"
	"flag"
	"time"
	"sync"
)

var verbose = flag.Bool("v", false, "show verbose progress message")

func main(){
	// 最初のディレクトリを決める
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	// ファイルツリーを走査する（並列）
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileSizes)
	}
	go func(){
		n.Wait()
		close(fileSizes)
	}()

	// 結果を表示する（-v オプションで定期的に表示する）
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

// sync.WaitGroup のコードを追加
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64){
	defer n.Done()
	for _, entry := range dirents(dir){
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents での並行性を制限するための計数セマフォ
var sema = make(chan struct{}, 20)

// sema による計数セマフォのコードを追加
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}
	defer func(){ <- sema }()

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du2: %v\n", err)
		return nil
	}
	return entries
}

// du1 の printDiskUsage と同じ。
func printDiskUsage(nfiles, nbytes int64){
	fmt.Printf("%d files  %1.f MB\n", nfiles, float64(nbytes)/1e6)
}
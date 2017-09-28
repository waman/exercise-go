// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// Netcat1 は読み込み専用の TCP クライアントです。
package main

import (
	"net"
	"log"
	"os"
	"io"
)

func main(){
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader){
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

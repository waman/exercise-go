// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// clock2 は時刻を定期的に書き出す TCP サーバです。
// 接続を並行して処理します。
package main

import (
	"net"
	"log"
	"io"
	"time"
)

// クライアントは ch8/netcat1 を使います。
func main(){
	log.Println("clock1 starts...")
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)  // 接続を並行して処理する
	}
}

func handleConn(c net.Conn){
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

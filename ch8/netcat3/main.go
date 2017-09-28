// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
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
	done := make(chan struct{})
	go func(){
		io.Copy(os.Stdout, conn)
		log.Println("done")
		done <- struct{}{}
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	<- done
}

// mustCopy は netcat1 のものと同じです。
func mustCopy(dst io.Writer, src io.Reader){
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
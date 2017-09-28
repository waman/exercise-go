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
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

// mustCopy は netcat1 のものと同じです。
func mustCopy(dst io.Writer, src io.Reader){
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
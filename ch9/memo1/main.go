package main

import (
	"github.com/waman/exercise-go/ch9/memo1/memo"
	"net/http"
	"io/ioutil"
	"time"
	"log"
	"fmt"
	"os"
)

func main(){
	m := memo.New(httpGetBody)
	for _, url := range os.Args[1:] {  // 本文から少し修正
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

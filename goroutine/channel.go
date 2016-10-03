package main

import (
	"net/http"
	"log"
	"fmt"
	"time"
	"runtime"
)

func main() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	// stringを扱うチャンネルを生成
	cpus := runtime.NumCPU()
	fmt.Println("cpus:", cpus)
	//statusChan := make(chan string)
	statusChan := make(chan string, cpus)

	for _, url := range urls {
		go func (url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			time.Sleep(15 * time.Second)
			defer res.Body.Close()

			// チャンネルに書き込む
			//statusChan <- res.Status
			statusChan <- url
		}(url)
	}

	// statusChanの読み出しが３回完了するまで処理がブロックされるためwaitGroup不要
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}

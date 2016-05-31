package main

import (
	"net/http"
	"log"
	"fmt"
)

/* channel.goからリクエストを行う処理をgetStatusに変更 */
// 戻り値である statusChanを「<-chan string」読みだし専用にした
func getStatus(urls []string) <- chan string {
	statusChan := make(chan string)

	for _, url := range urls {
		go func (url string) {
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			}
			defer res.Body.Close()
			//statusChan <- res.Status
			statusChan <- url
		}(url)
	}
	return statusChan
}

func main() {
	urls := []string{
		"http://example.com",
		"http://example.net",
		"http://example.org",
	}

	// getStatusで受け取ったチャンネルをstatusChan変数に読み出す
	statusChan := getStatus(urls)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-statusChan)
	}
}

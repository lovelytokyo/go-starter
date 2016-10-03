package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("start")
	var wg sync.WaitGroup
	for _, a := range []string{"1", "2"} {
		wg.Add(1)
		go func(str string) {
			i := 0
			fmt.Println("I'm " + str)
			time.Sleep(1) // 時分割なので2番目が開始するまで待つ
			for {
				i++;
			}
			wg.Done()
		}(a)
	}
	wg.Wait()
	fmt.Println("end")
}
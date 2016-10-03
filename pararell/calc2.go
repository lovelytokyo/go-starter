package main

import (
	"fmt"
	"sync"
	"time"
	"runtime"
)

func main() {
	fmt.Println("start")
	fmt.Printf("NumCPU=%d\n", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())

	var wg sync.WaitGroup
	for _, a := range []string{"1", "2"} {
		wg.Add(1)
		go func(str string) {
			i := 0
			fmt.Println("I'm " + str)
			time.Sleep(1)
			for {
				i++;
			}
			wg.Done()
		}(a)
	}
	wg.Wait()
	fmt.Println("end")
}
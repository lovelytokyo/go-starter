package main

import (
	"fmt"
	"calc" // $GOPATH/src 以下のディレクトリ構造を表した形でimport
)

func main() {
	fmt.Printf("Add(%v, %v) = %v\n", 1, 2, calc.Add(1,2))
}

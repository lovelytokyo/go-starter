package main

import (
	"fmt"
)

func swap(i, j int) (int, int) {
	return j, i
}

func main() {
	x, y := 3, 4
	x, y = swap(x, y)
	fmt.Println(x, y) // 4 3

	// x = swap(x, y) // compile error
	x, _ = swap(x, y) // 第２戻り値を無視
	fmt.Println(x, y) // 3 3

	swap(x, y) // compile、実行ともに可能
	fmt.Println(x, y) // 3 3
}

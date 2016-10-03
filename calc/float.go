package main

import "fmt"

func main() {
	var a float32 = 0.1234
	var b float64 = 0.5678

	fmt.Print("(1) a = ")
	fmt.Printf("%f", a)
	fmt.Print("\n")

	// 以下の書き方が一般的
	fmt.Printf("(2) a = %6.2f\n", a)
	fmt.Printf("(3) b = %6.3f\n", b)
}
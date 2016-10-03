package main

import (
	"strconv"
	"fmt"
)

func main() {
	click := 30000000
	imp := 30490
	ctr := fmt.Sprintf("%0.2f", float64(click) / float64(imp) * 100)
	ctr2 := fmt.Sprintf("%4.2f", float64(click) / float64(imp) * 100)
	fmt.Printf("ctr: %v", ctr)
	fmt.Println("")
	fmt.Printf("ctr2: %v", ctr2)
	fmt.Println("")

	f, err := strconv.ParseFloat(ctr, 64)
	fmt.Printf("f: %v", f)
	fmt.Println("")
	fmt.Printf("err: %v", err)
	fmt.Println("")

	f2, err := strconv.ParseFloat(ctr2, 64)
	fmt.Printf("f: %v", f2)
	fmt.Println("")
	fmt.Printf("err: %v", err)
	fmt.Println("")

	ctvr := 98432.412000000000006
	fmt.Printf("%0.3f", ctvr)
	fmt.Println("")
	parsed, err := strconv.ParseFloat(fmt.Sprintf("%0.3f", ctvr), 64)

	fmt.Print(parsed)
}

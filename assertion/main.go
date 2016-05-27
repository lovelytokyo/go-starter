package main

import "fmt"

func Print(value interface{}) {
	s, ok := value.(string) //type assertion
	if ok {
		fmt.Printf("value is string: %s\n", s)
	} else {
		fmt.Println("value is not string \n")
	}
}

func main() {
	Print("abc")
	Print(1)
}

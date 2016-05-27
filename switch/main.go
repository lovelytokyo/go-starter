package main

import "fmt"

type Stringer interface {
	String() string
}

func Print(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("value is string: %s\n", v)
	case int:
		fmt.Printf("value is int: %d\n", v)
	case Stringer:
		fmt.Printf("value is Stringer: %s\n", v)
	}
}

func main() {
	Print("abc")
	Print(10)
}

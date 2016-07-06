package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) ChangeNameUnsafe(name string) {
	p.name = name
}

func (p *Person) ChangeNameSafe(name string) {
	if p == nil {
		return
	}
	fmt.Println("name:", name)
	p.name = name
}

func main() {
	var pp *Person // nil ポインタ変数
	// pp.ChangeNameUnsafe("Jiro")  // => panic: runtime error: invalid memory address or nil pointer dereference
	pp.ChangeNameSafe("Jiro") //
	fmt.Printf("%v", pp)
}

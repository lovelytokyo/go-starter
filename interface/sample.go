package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human //an anonymous field of type Human
	school string
	loan float32
}

type Employee struct {
	Human //an anonymous field of type Human
	company string
	money float32
}

//A human method to say hi
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//A human can sing a song
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee's method overrides Human's one
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}

// INTERFACES
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{
		Human{
			"Mike", 25, "222-22-XXX",
		},
		"MIT",
		0.00,
	}

	paul := Student{
		Human{
			"Paul", 26, "111-222-XXX",
		},
		"Harvard",
		100,
	}

	sam := Employee{
		Human{
			"Sam", 36, "434-234-3434",
		},
		"Golan Inc",
		1000,
	}

	Tom := Employee{
		Human{
			"Sam", 36, "434-234-3434",
		},
		"Things Ltd.",
		5000,
	}

	// a variable of the interface type Men
	var i Men
	i = mike
	fmt.Println("This is Mike, a Student")
	i.SayHi()
	i.Sing("November rain")

	i = Tom
	fmt.Println("This is Tom, an Employee")
	i.SayHi()
	i.Sing("Born to be wild")

	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	x[0], x[1], x[2] = paul, sam, mike

	for _, value := range x{
		value.SayHi()
	}
}
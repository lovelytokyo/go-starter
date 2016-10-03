package main

import (
	"fmt"
	"gopkg.in/alioygur/godash.v0"
)

// image size
const (
	TYPE300  = "300×300"
	TYPE1200 = "1200×628"
)

var imageTypes = map[string]string{
	TYPE300:  "300_300",
	TYPE1200: "1200_628",
}

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
	//Print("abc")
	//Print(10)

	var s string
	s = "300_200"
	if godash.IsNull(s) {
		fmt.Printf("null: %s", imageTypes[TYPE300])
		return
	}

	//imageType, ok := imageTypes[s]
	//if !ok {
	//	fmt.Printf("not setted: %s", imageTypes[TYPE300])
	//	return
	//}
	//
	//fmt.Printf("setted: %s", imageType)
	//returns

	switch s {
	case TYPE300:
		fmt.Printf("300: %s",  imageTypes[s])
	case TYPE1200:
		fmt.Printf("1200: %s",  imageTypes[s])
	default:
		fmt.Printf("default: %s",  imageTypes[TYPE300])
	}
}

package main

import (
	"fmt"
	"reflect"
)

func call(f interface{}) {
	fv := reflect.ValueOf(f)
	if fv.Kind() != reflect.Func {
		panic("f must be func.")
	}
	fv.Call([]reflect.Value{})
}

func main() {
	f := func() {
		fmt.Println("hello!")
	}
	call(f)
}

package main

import (
	"fmt"
	"reflect"
)

func set(p, v interface{}) error {
	pv := reflect.ValueOf(p)
	if pv.Kind() != reflect.Ptr {
		return fmt.Errorf("p must be pointer.")
	}

	vv := reflect.ValueOf(v)
	if pv.Elem().Kind() != vv.Kind() {
		return fmt.Errorf("p type and v type do not mutch")
	}

	pv.Elem().Set(vv)

	return nil
}

func main() {
	var hoge int
	fmt.Println(hoge)
	set(&hoge, 100)
	fmt.Println(hoge)
	fmt.Println(set(&hoge, 10.4))
}
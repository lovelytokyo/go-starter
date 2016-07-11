package main

import "fmt"
import "reflect"

func AbstractSlice(arr interface{}) []interface{} {
	dest := []interface{}{}
	switch sl := arr.(type) {
	case []interface{}:
	// こういうのは無理なんですよ
	case string:
		// rangeかけるためにcase内じゃないとダメ
		for _, b := range sl {
			dest = append(dest, b)
		}
	case []string:
		for _, str := range sl {
			dest = append(dest, str)
		}
	case []int:
		for _, i := range sl {
			dest = append(dest, i)
		}
	}
	return dest
}

func main() {

	a := "slice of byte"
	fmt.Println(
		reflect.TypeOf(a),
		reflect.TypeOf(AbstractSlice(a)),
	)

	b := []string{"Hello", "interface"}
	fmt.Println(
		reflect.TypeOf(b),
		reflect.TypeOf(AbstractSlice(b)),
	)

	c := []int{2, 3, 4, 5}
	fmt.Println(
		reflect.TypeOf(c),
		reflect.TypeOf(AbstractSlice(c)),
	)
}
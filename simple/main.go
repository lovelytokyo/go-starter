package main

import(
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	json := simplejson.New()
	json.Set("message", "Hello World!")

	b, _ := json.EncodePretty()

	fmt.Printf("%s\n", b)
}

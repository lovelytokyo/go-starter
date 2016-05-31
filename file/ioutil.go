package main

import (
	"io/ioutil"
	"log"
	"fmt"
)

func main()  {

	/* write file some message*/
	hello := []byte("hello world My Go Stater\n")
	err := ioutil.WriteFile("./file.txt", hello, 0666)
	if err != nil {
		log.Fatal(err)
	}

	/* read file */
	message, err := ioutil.ReadFile("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(message))
}

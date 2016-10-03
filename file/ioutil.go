package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"strings"
)

func main()  {

	s:= []string{
		"hello world Taro",
		"hello world Michi",
	}
	t := strings.Join(s, "\n")

	/* write file some message*/
	//hello := []byte("hello world My Go Stater\n")
	hello := []byte(t)
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

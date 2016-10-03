package main

import (
	"strings"
	"github.com/k0kubun/pp"
)

func main() {
	//ids := ""
	ids := "test1,test2"

	var n int
	result := strings.SplitN(ids, ",", n)
	pp.Println(result)
	pp.Println(n)
}




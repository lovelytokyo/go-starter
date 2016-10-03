package main

import (
	"flag"
	"fmt"
)

func main() {
	var ip *int = flag.Int("flagname", 1234, "help message for flagname")
	var showVersion *bool = flag.Bool("v", false, "show version")
	var version string = "1.35"

	flag.Parse()
	fmt.Println(*ip)
	fmt.Println(*showVersion)

	if *showVersion {
		fmt.Printf("version: %s\n", version)
		return
	}

	fmt.Println("end")
}

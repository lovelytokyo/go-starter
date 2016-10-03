package main

import (
	"os/exec"
	"fmt"
)

func printErr(err error) {
	fmt.Println(err)
	return
}

func main() {
	date, err := exec.Command("date").Output()
	if err != nil {
		printErr(err)
	}
	fmt.Printf("The date is %s\n", date)

	out, err := exec.Command("pwd").Output()
	if err != nil {
		printErr(err)
		return
	}
	fmt.Printf("pwd is %s\n", out)


	output, err := exec.Command("ls", "-la").Output()
	if err != nil {
		printErr(err)
		return
	}
	fmt.Printf("list is %s\n", output)
}

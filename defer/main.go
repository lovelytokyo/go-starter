package main

import (
	"os"
	"fmt"
	"errors"
)

func main() {
	file, err := os.Open("./error.go")
	if err != nil {
		// エラー処理
		fmt.Println("error processing")
	}

	fmt.Println("file exists")
	errors.New("error something")

	// 関数を抜ける前にかならず実行される
	defer file.Close()
	// 正常処理

}

package main

import (
	"os"
	"log"
	"fmt"
)

func main() {
	file, err := os.Create("./file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 書き込んだデータを[]byteで用意する
	message := []byte("hello world\n")

	// Write関数は、書き込んだバイト数とエラーを返す
	_, err = file.Write(message)
	if err != nil {
		log.Fatal(err)
	}

	file2, err2 := os.Create("./file2.txt")
	_, err2 = file2.WriteString("hello world2\n")
	if err2 != nil {
		log.Fatal(err2)
	}

	/* ファイル読み込み */
	file, err3 := os.Open("./file.txt")
	if err3 != nil {
		log.Fatal(err3)
	}

	// 12byte格納可能なスライスを用意する
	msg := make([]byte, 12)

	// ファイル内のデータをスライスに読み出す
	_, err = file.Read(msg)
	if err != nil {
		log.Fatal(err)
	}

	// 文字列にして表示
	fmt.Println(string(msg))
}

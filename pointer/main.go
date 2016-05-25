package main

import "fmt"

func callByValue(i int) {
	i = 20 // 上書き
}

func callByRef (i *int) { // ポインター型変数
	*i = 20 // 参照先を上書きする
}

func main() {
	var i int = 10
	callByValue(i) // 値を渡す
	fmt.Println(i) // 10
	callByRef(&i) // 変数のアドレスを渡す
	fmt.Println(i) // 20
}

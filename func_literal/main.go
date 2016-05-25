package main;

import(
	"fmt"
)

/*
var sum func(i, j int) = func (i, j int) {
	fmt.Println(i + j)
}
*/

func main() {
	//sum(2, 4)

	// 無名関数定義と即時実行
	func(i, j int) {
		fmt.Println(i + j)
	} (2, 4)
}

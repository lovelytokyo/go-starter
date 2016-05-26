package main

import (
	"log"
	"fmt"
	"errors"
)

func main() {
	defer func() { // 関数を抜ける前にかならず実行される
		err := recover()
		if err != nil {
			// runtime error : index out or range
			log.Fatal(err)
		}
	} ()

	a := []int{1,2,3}
	fmt.Println(a[10]) // occered paninc

	/* 実行結果 */
	//2016/05/26 09:40:17 runtime error: index out of range
	//exit status 1

	/* 自分でpanic発生される例 */
	// panicを用いるのは、エラーを戻り値として表現できない場合や回復が不可能なシステムエラー、やむを得つ大域脱出が必要な場合など
	// 基本的にエラーは、関数の戻り値として呼び出し側に返すべき
	b := []int{1, 2, 3}
	for i := 0; i < 10; i++ {
		if i >= len(b) {
			panic(errors.New("Index out of range"))
		}
		fmt.Println(b[i])
	}

}

package main

import (
	"errors"
	"fmt"
	"log"
)

/*
関数が多値を返せることを利用して，内部で発生したエラーを戻り値で表現します。
関数の処理に成功した場合はエラーはnilにし，
異常があった場合はエラーだけに値が入り，他方はゼロ値になります。
 */
func div(i, j int) (int, error) {
	if j == 0 {
		// 自作エラーを返す
		return 0, errors.New("divied by zero")
	}
	return i / j, nil
}

/* 名前付き返り値 */
func div2(i, j int) (result int, err error) {
	if j == 0 {
		err = errors.New("divied by zero")
		return // return 0, erroと同じ
	}
	result = i / j
	return; // return result, nilと同じ
}

func main() {
	n, err := div2(10, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

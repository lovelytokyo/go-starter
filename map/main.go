package main

import "fmt"

func main()  {
	/* 宣言と初期化 */
	// 宣言
	var month map[int]string = map[int]string{}

	// 初期化
	month[1] = "January"
	month[2] = "Fabruary"
	fmt.Println(month)

	/* 宣言と初期化を同時に行う */
	month2 := map[int]string {
		1: "January",
		2: "Fabruary",
	}
	fmt.Println(month2)

	/* mapの操作 */
	// データ取得
	//jan := month[1]
	//_, ok := month[1]  // マップ内のキー存在を調べるのみの場合は、値を無視する「_」
	jan, ok := month[1]
	fmt.Println(ok)
	if ok {
		fmt.Println("map操作")
		fmt.Println(jan)
	}

	// データ削除
	delete(month, 1)
	fmt.Println(month)

	month[1] = "January"
	month[3] = "March"
	month[4] = "April"

	/* range */
	// スライスとは違い、順番保証されない
	fmt.Println("range")
	for key, value := range month {
		fmt.Printf("%d %s\n", key, value)
	}

}

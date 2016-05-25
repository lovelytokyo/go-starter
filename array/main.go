package main

import "fmt"

func fn(arr [4]string) {
	arr[0] = "x" // main関数には反映されない
	fmt.Println(arr) // [x b c d]
}

func sum(nums ...int) (result int) {
	// numsは[]int型
	for _, n := range nums {
		result += n
	}
	return
}

func main() {
	/* 配列定義、初期化方法  */
	// 配列の定義、初期化方法１
	//var arr [4]string
	//arr[0] = "a"
	//arr[1] = "b"
	//arr[2] = "c"
	//arr[3] = "d"

	// 配列の定義、初期化方法２
	// arr := [4]string{"a", "b", "c", "d"}
	 arr := [...]string{"a", "b", "c", "d"}

	i := 0
	for(i < len(arr)) {
		fmt.Println(arr[i])
		i++
	}

	/* 引数の型 */
	arr1 := [4]string{"a", "b", "c", "d"}
	//arr2 := [5]string{"a", "b", "c", "d", "e"}
	fn(arr1)
	//fn(arr2) // arr1とarr2の要素の型はStringですが、関数fnは[4]string型を引数ととるためコンパイルエラーとなる

	/* 関数に渡したパラメータが配列の場合は、配列のコピーが渡される */
	fmt.Println(arr1) // [a b c d]

	/* append */
	var s []string // スライスの宣言。配列のように長さ情報がない
	s = append(s, "a")
	s = append(s, "b")
	s = append(s, "c", "d")
	fmt.Printf("append : %s\n" , s)

	/* range */
	fmt.Println("range")
	for i, s := range arr1 {
		// i = 添字 s = 値
		fmt.Println(i, s)
	}

	/* 値の切り出し */
	slice := []int{0,1,2,3,4,5}
	fmt.Println("slice")
	fmt.Println(slice[2:4]) // 2 3
	fmt.Println(slice[0:len(slice)]) // 0 1 2 3 4 5
	fmt.Println(slice[:3]) // 0 1 2
	fmt.Println(slice[3:]) // 3 4 5
	fmt.Println(slice[:]) // 0 1 2 3 4 5

	/* 可変長引数 */
	fmt.Printf("可変長引数: %d\n", sum(1,2,3,4))

}

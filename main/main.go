package main
import (
	"fmt"
	"gosample"
)

func main() {
	fmt.Println(gosample.Message) // hello world

	// 初期値
	var i int;
	var j string;
	var k bool;

	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(nil)

	// if
	a, b := 10, 100
	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a = b")
	}

	// for
	n := 0;
	for (n < 10) {
		fmt.Printf("n = %d\n", n);
		n++
	}

	// break, conntinue
	x := 0
	for { // 無限ループ
		x++
		if x > 10 {
			break // ループ抜ける
		}
		if x % 2 == 0 {
			continue // 偶数は飛ばす
		}
		fmt.Println(x) // 奇数のみ表示
	}

	// switch
	q := 10
	switch q {
	case 15:
		fmt.Println("FizzBuzz")
	case 5, 10:
		fmt.Println("Buzz")
	case 3, 6, 9:
		fmt.Println("Fizz")
	default:
		fmt.Println(q)
	}

	// 順次caseに処理進めたい場合
	ft := 3
	switch ft {
	case 3:
		ft = ft -1
		fallthrough
	case 2:
		ft = ft -1
		fallthrough
	case 1:
		ft = ft -1
		fmt.Println(ft)
	}

	xo := 10
	switch  {
	case xo%15 == 0:
		fmt.Println("FizzBuzz")
	case xo%5 == 0:
		fmt.Println("Buzz")
	case xo%3 == 0:
		fmt.Println("Fizz")
	default:
		fmt.Println(xo)
	}
}
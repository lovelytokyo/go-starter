package main

import "fmt"

type Person struct {
	Name string
}

// Person 型に対してメソッドを定義する : 値レシーバ
func (p Person) Greet(msg string) {
	fmt.Printf("%s, I'm %s. \n", msg, p.Name)
}

// *Person 型に対してメソッドを定義する : ポインタレシーバ
func (pp *Person) Shout(msg string) {
	fmt.Printf("%s!!!  \n", msg)

	// フィールド（pp.Name）を使うとポインタ変数から呼ばれた祭にpanicになる
	// fmt.Printf("%s!!! %s! \n", msg, pp.Name)
}

func main() {
	p := Person {Name: "Taro"} // 値型の変数を用意する
	p.Greet("Hi")		// => Hi, I'm Taro.

	pp := &Person{Name: "Taro"} // ポインタ型の変数を用意する
	pp.Shout("OH MY GOD") // => OH MY GOD!!! Taro!

	(*pp).Greet("Hi") // => Hi, I'm Taro || 当然呼び出せる
	pp.Greet("Hi") // => Hi, I'm Taro  || コンパイラーが↑のコードに変換してくれる

	// 値レシーバの場合、ポインタ変数から呼び出せない
	//var nilp *Person // nilポインタ変数
	//nilp.Greet("Hi") // => panic  runtime error: invalid memory address or nil pointer dereference

	// 値型変数のレシーバを呼ぶ
	(&p).Shout("OH MY GOD") // => OH MY GOD!!! Taro! || 当然呼び出せる
	p.Shout("OH MY GOD") // => OH MY GOD!!! Taro! || コンパイラが↑のコードに変換してくれる

	// ポインタレシーバの場合、nilポインタ変数から呼び出し可能
	var nilp *Person
	nilp.Shout("OH MY GOD") // => OH MY GOD!!!  || 呼び出せる。ただし、メソッド内でフィールドを使ってたら参照先がないためpanicになる
}

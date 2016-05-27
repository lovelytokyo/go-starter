package main

import "fmt"

// 構造体定義
type Task struct {
	ID int
	Detail string // 大文字で始まる場合はPublic
	done bool // 小文字は、パッケージ内に閉じたスコープ
}

type Task2 struct {
	ID int
	Detail string // 大文字で始まる場合はPublic
	done bool // 小文字は、パッケージ内に閉じたスコープ
}

func main()  {
	// 構造体生成
	var task Task = Task {
		ID : 1,
		Detail : "Buy the milk",
		done : true,
	}
	fmt.Println(task.ID)
	fmt.Println(task.Detail)
	fmt.Println(task.done)

	// フィールド名を省略
	var task2 Task2 = Task2 {
		1, "Buy the bread", true,
	}

	fmt.Println(task2.ID)
	fmt.Println(task2.Detail)
	fmt.Println(task2.done)

	// 構造体生成時、値はゼロ値で初期化される
	task3 := Task{}
	fmt.Println(task3.ID)
	fmt.Println(task3.Detail)
	fmt.Println(task3.done)

}

package main

import "fmt"

type ID int  // type 名前 型。int型を拡張したID型を定義
type Priority int

func ProcessTask (id ID, pr Priority) {
	fmt.Printf("id: %d, priority: %d\n", id, pr)
}

func ProcessTask2 (id2 int , priority2 int) {
	fmt.Printf("id: %d, priority: %d", id2, priority2)
}

func main()  {
	var id ID = 3
	var pr Priority = 5
	ProcessTask(id, pr)

	/* type宣言していない場合は、パラメータ順番が変わってもすぐに検知することができない */
	id2 := 3;
	priority2 := 5;
	ProcessTask2(priority2, id2)
}



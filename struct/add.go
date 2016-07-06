package main

import "fmt"

type MyBoxItem struct {
	Name string
}

type MyBox struct {
	Items []MyBoxItem
}

func (box *MyBox) AddItem(item MyBoxItem) []MyBoxItem {
	box.Items = append(box.Items, item)
	return box.Items
}

func main() {
	item1 := MyBoxItem{Name: "Test Item 1"}

	items := []MyBoxItem{}
	box := MyBox{items}
	box.AddItem(item1)
	fmt.Println(box.Items) // => [{Test Item 1}]
}

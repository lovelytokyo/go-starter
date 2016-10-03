package main

import "fmt"
import "sort"
func main() {
	m := map[string]int{"ada": 1, "hoge": 4, "basha": 3, "poeni": 3}

	a := List{}
	for k, v := range m {
		e := Entry{k, v}
		a = append(a, e)
	}

	sort.Sort(a)
	fmt.Println(a)
}

type Entry struct {
	name  string
	value int
}
type List []Entry

func (l List) Len() int {
	return len(l)
}

func (l List) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func (l List) Less(i, j int) bool {
	if l[i].value == l[j].value {
		return (l[i].name < l[j].name)
	} else {
		return (l[i].value < l[j].value)
	}
}
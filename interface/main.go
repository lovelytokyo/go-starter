package main

import "fmt"

type Task struct {
	ID int
	Detail string
	done bool
}

/* ToString関数のinterface */
type ToStringer interface {
	ToString() string
}

func NewTask(id int, detail string) *Task {
	task := &Task {
		ID: id,
		Detail: detail,
		done: false,
	}
	return task
}

/* Taskの文字列表現を返すStringメソッド */
func (task Task) ToString() string {
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

func Print(toStringer ToStringer) {
	fmt.Println(toStringer.ToString())
}

func main() {
	task := NewTask(1, "Buy the Milk")
	Print(task)
}



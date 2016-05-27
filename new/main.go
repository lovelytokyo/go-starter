package main

import "fmt"

type Task struct {
	ID int
	Detail string
	done bool
}

func main() {
	var task *Task = new(Task)
	task.ID = 1
	task.Detail = "buy the milk"
	fmt.Println(task.done) // false
}

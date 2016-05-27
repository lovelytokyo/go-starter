package main

import "fmt"

type Task struct {
	ID int
	Detail string
	done bool
}

type Todo struct {
	ID int
	Detail string
	done bool
}


func Finish(task Task, todo *Todo) {
	task.done = true
	todo.done = true
}

func main() {
	task := Task{done: false}
	todo := &Todo{done: false}

	Finish(task, todo)

	fmt.Println(task.done) //false
	fmt.Println(todo.done) // true
}

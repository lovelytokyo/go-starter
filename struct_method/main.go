package main

import "fmt"

type Task struct {
	ID int
	Detail string
	done bool
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
	fmt.Println("ToString")
	// レシーバ「task」はTaskのコピー渡されるため、値変更できない
	task.Detail = "buy the bread"
	str := fmt.Sprintf("%d) %s", task.ID, task.Detail)
	return str
}

/* レシーバをポインターとして受け取って、タスクを終了済みにするFinishメソッド */
func (task *Task) Finish() {
	fmt.Println("Finish")
	task.done = true
}

func main() {
	task := NewTask(1, "buy the milk")
	fmt.Printf("%+v\n", task)
	// &{ID:1 Detail:buy the milk done:false}

	fmt.Println(task.ToString())
	// 1) buy the bread

	fmt.Printf("%+v\n", task) // task.DetailがToString内部で変更されているが、反映されていない
	// &{ID:1 Detail:buy the milk done:false}


	task.Finish()
	fmt.Printf("%+v\n", task) // task.doneがFinishで変更されているのが反映されている
	// &{ID:1 Detail:buy the milk done:true}

}

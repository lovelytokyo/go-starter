package main

import "fmt"

type Task struct {
	ID int
	Detail string
	done bool
	*User // Userを埋め込む
}

type User struct {
	FirstName string
	LastName string
}

func NewTask(id int, detail , firstName, lastName string) *Task {
	fmt.Printf("User : %+v\n", NewUser(firstName, lastName))
	task := &Task {
		ID: id,
		Detail: detail,
		done: false,
		User: NewUser(firstName, lastName),
	}
	fmt.Printf("task : %+v\n", task)
	return task
}

func NewUser(firstName, lastName string) *User {
	return &User {
		FirstName: firstName,
		LastName : lastName,
	}
}

func (user *User) FullName() string {
	fullname := fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	return fullname
}




func main() {
	// TaskにUserフィールドが埋め込まれている
	task := NewTask(1, "buy the milk", "Jackson", "Michel")
	fmt.Printf("%+v", task)

	// User.FirstName, User.LastName、Usesr.fullName()がTaskが実装しているかのように振る舞う
	fmt.Println("firstName : " + task.FirstName)
	fmt.Println("lastName : " + task.LastName)
	fmt.Println("fullName : " + task.FullName())
	fmt.Println(task.User)
}

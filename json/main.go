package main

import (
	"encoding/json"
	"log"
	"fmt"
)

type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"-"`
	memo string
}

func main() {
	// 構造体をjsonへ変換
	person := &Person {
		ID: 1,
		Name: "Gopher",
		Email: "gopher@example.org",
		memo : "golang lover",
	}

	b, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))

	// jsonを構造体に変換
	var person2 Person
	c := []byte(`{"id":2,"name":"Janifer","email":"janifer@example.org"}`)
	err2 := json.Unmarshal(c, &person2)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(person2)
}

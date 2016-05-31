package main

import (
	"encoding/json"
	"log"
	"fmt"
	"os"
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

	/* write json to file by json encoder */
	// open file
	file, err := os.Create("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// get encoder
	encoder := json.NewEncoder(file)

	// encode json and write to file
	err = encoder.Encode(person)
	if err != nil {
		log.Fatal(err)
	}
	//$ cat person.json                                                                                                                                                    (master✱)
	//{"id":1,"name":"Gopher"}


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

	/* read file and decode string to json */
	file3, err := os.Open("./person.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file3.Close()

	var person3 Person

	// get data
	decoder := json.NewDecoder(file3)

	// write json data to type
	err = decoder.Decode(&person3)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(person3)
}

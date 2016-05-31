package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"log"
	"os"
	"html/template"
	"strconv"
	"io/ioutil"
)

type Person struct {
	ID int `json:"id"`
	Name string `json:"name"`
}

var t = template.Must(template.ParseFiles("index.html"))

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "hello world")
}

func PersonHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close() // 処理の最後にBodyを閉じる

	// decode response body to json
	if r.Method == "POST" {
		var person Person
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}

		// set filename {id}.txt
		filename := fmt.Sprintf("%d.txt", person.ID)
		file, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// write name to file
		_, err = file.WriteString(person.Name)
		if err != nil {
			log.Fatal(err)
		}

		// send response stasus code 201
		w.WriteHeader(http.StatusCreated)
	} else if r.Method == "GET" {

		// get parameter
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("%d.txt", id)
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Fatal(err)
		}

		// create person
		person := Person{
			ID: id,
			Name: string(b),
		}

		// write encoded html to response
		t.Execute(w, person)
	}

}

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/persons", PersonHandler)
	http.ListenAndServe(":3000", nil)
}

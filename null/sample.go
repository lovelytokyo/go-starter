package main

import (
	null "gopkg.in/guregu/null.v3"
	zero "gopkg.in/guregu/null.v3/zero"
	"database/sql"
	"encoding/json"
	"fmt"
)

type DBUser struct {
	Id uint32 `json:"id"`
	FirstName string `json:"first_name"`
	MiddleName sql.NullString `json:"middle_name"`
	LastName string `json:"last_name"`
	Salary sql.NullInt64 `json:"salary"`
}

type JsonNullUser struct {
	Id uint32 `json:"id"`
	FirstName string `json:"first_name"`
	MiddleName null.String `json:"middle_name"`
	LastName string `json:"last_name"`
	Salary null.Int `json:"salary"`
}

type JsonZeroUser struct {
	Id uint32 `json:"id"`
	FirstName string `json:"first_name"`
	MiddleName zero.String `json:"middle_name"`
	LastName string `json:"last_name"`
	Salary zero.Int `json:"salary"`
}

func main() {

	// --- sql.nullString
	user1 := &DBUser {
		Id: 1,
		FirstName: "Minjung",
		LastName: "Oh",
	}
	user1Json, _ := json.Marshal(user1)
	fmt.Println(string(user1Json))
	// {"id":1,"first_name":"Minjung","middle_name":{"String":"","Valid":false},"last_name":"Oh","salary":{"Int64":0,"Valid":false}}


	// --- null.String, null.Int
	user2 := &JsonNullUser{
		Id: 2,
		FirstName: "Mark",
		LastName: "Parker",
	}

	user2Json, _ := json.Marshal(user2)
	fmt.Println(string(user2Json))
	// {"id":2,"first_name":"Mark","middle_name":null,"last_name":"Parker","salary":null}

	user21 := &JsonNullUser{
		Id: 21,
		FirstName: "Mark",
		MiddleName: null.NewString("", true),
		LastName: "Parker",
		Salary: null.NewInt(0, true),
	}

	user21Json, _ := json.Marshal(user21)
	fmt.Println(string(user21Json))
	// {"id":21,"first_name":"Mark","middle_name":"","last_name":"Parker","salary":0}

	user3 := &JsonNullUser{
		Id: 3,
		FirstName: "Jack",
		MiddleName: null.NewString("Junior", true),
		LastName: "Paul",
		Salary: null.NewInt(30, true),
	}

	user3Json, _ := json.Marshal(user3)
	fmt.Println(string(user3Json))
	// {"id":3,"first_name":"Jack","middle_name":"Junior","last_name":"Paul","salary":30}

	// --- zero.String, zeroInt
	user4 := &JsonZeroUser{
		Id: 4,
		FirstName: "Mark",
		LastName: "Parker",
	}
	user4Json, _ := json.Marshal(user4)
	fmt.Println(string(user4Json))
	// {"id":4,"first_name":"Mark","middle_name":"","last_name":"Parker","salary":0}

	user41 := &JsonZeroUser{
		Id: 41,
		FirstName: "Mark",
		MiddleName: zero.NewString("", true),
		LastName: "Parker",
		Salary: zero.NewInt(0, true),
	}
	user41Json, _ := json.Marshal(user41)
	fmt.Println(string(user41Json))
	// {"id":41,"first_name":"Mark","middle_name":"","last_name":"Parker","salary":0}

	user5 := &JsonZeroUser{
		Id: 5,
		FirstName: "Jack",
		MiddleName: zero.NewString("Junior", true),
		LastName: "Paul",
		Salary: zero.NewInt(30, true),
	}

	user5Json, _ := json.Marshal(user5)
	fmt.Println(string(user5Json))
	// {"id":5,"first_name":"Jack","middle_name":"Junior","last_name":"Paul","salary":30}
}

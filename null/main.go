package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
)

type User struct {
	Id uint32 `json:"id"`
	FirstName string `json:"first_name"`
	MiddleName sql.NullString `json:"middle_name"`
	LastName string `json:"last_name"`
}

func main() {
	var tmp sql.NullString
	fmt.Printf("%T(%+v)", tmp, tmp)
	// sql.NullString({String: Valid:false})

	user1 := &User {
		Id: 1,
		FirstName: "Minjung",
		LastName: "Oh",
	}
	user1Json, _ := json.Marshal(user1)
	fmt.Println(string(user1Json))
	// {"id":1,"first_name":"Minjung","middle_name":{"String":"","Valid":false},"last_name":"Oh"}

	user2 := &User {
		Id: 2,
		FirstName: "Mishel",
		MiddleName: sql.NullString {
			"Junior",
			true,
		},
		LastName: "Tom",
	}
	user2Json, _ := json.Marshal(user2)
	fmt.Println(string(user2Json))
	// {"id":2,"first_name":"Mishel","middle_name":{"String":"Junior","Valid":true},"last_name":"Tom"}

	user3 := &User {
		Id: 2,
		FirstName: "Peter",
		MiddleName: sql.NullString {
			"",
			false,
		},
		LastName: "Jacson",
	}

	uset3Json, _ := user3.MarshalJSON()
	fmt.Printf("%+v", string(uset3Json))
}

func (u *User) MarshalJSON() ([]byte, error){
	middleNameValue, err := u.MiddleName.Value()

	if err != nil {
		return nil, err
	}

	var middleNameJsonString string

	if middleNameValue == nil {
		middleNameJsonString = "null"
	} else {
		middleNameJsonString = fmt.Sprintf("\"%s\"", middleNameValue)
	}

	jsonString := fmt.Sprintf("{\"id\":%d,\"first_name\":%s,\"middle_name\":%s,\"last_name\":%s}", u.Id, u.FirstName, middleNameJsonString, u.LastName)

	return []byte(jsonString), nil
}

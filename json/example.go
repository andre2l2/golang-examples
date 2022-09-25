package jsonexample

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Status bool
} 

func Example1() {
	user := User{"Jose", true}

	res, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(res))
}


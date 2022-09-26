package jsonexample

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Data struct {
	Name string
	Status bool
}

func Example2() {
	content, err := ioutil.ReadFile("./docs/user.json")

	if err != nil {
		log.Fatal("Error when ReadFile()", err)
	}

	var payload Data
	err = json.Unmarshal(content, &payload)

	if err != nil {
		log.Fatal("Error when Unmarshal()", err)
	}

	fmt.Println(payload)
}
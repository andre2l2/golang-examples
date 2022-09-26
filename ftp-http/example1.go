package ftp_http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status int
}

func Example1() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response, err := json.Marshal(Response{200})
		if err != nil {
			panic(err)
		}

		fmt.Fprint(w, string(response))
	})

	fmt.Println("Servcer runnug at 5050")
	http.ListenAndServe(":5050", nil)
}
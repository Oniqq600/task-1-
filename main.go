package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "hello, ", task)
	} else {
		fmt.Println("Ni ni ni")
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "Invalid request method")
	} else {
		fmt.Println("Всё ок")
	}

	var reqBody requestBody

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		fmt.Println("Error")
	}

	task = reqBody.Message

	fmt.Fprintln(w, "Обновлено на: ", task)
}

func main() {
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe("localhost:8080", nil)

}

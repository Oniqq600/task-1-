package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Message struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Я не знаю что за фатал!!!", err)
	}

	err = db.AutoMigrate(&Message{})
	if err != nil {
		log.Fatal("Ошибка при миграции: ", err)
	}
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Message
	w.Header().Set("Content-Type", "application/json")
	err := db.Find(&tasks)
	if err != nil {
		fmt.Println("Всё ещё для того что бы не светило")
	}

	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintln(w, "Я не знаю нахрена здесь в начале буква W")
	}

	var newTask Message
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		fmt.Println("Это что бы крвсным не горело")
	}

	db.Create(&newTask)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	InitDB()
	http.HandleFunc("/get", GetHandler)
	http.HandleFunc("/post", PostHandler)
	http.ListenAndServe("localhost:8080", nil)

}

package main

import (
	//"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Message struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
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

func GetHandler(e echo.Context) error {
	var messages []Message

	newTable := db.Find(&messages)
	if newTable == nil {
		return e.JSON(http.StatusOK, &messages)
	}

	return e.JSON(http.StatusOK, &messages)
}

func PostHandler(e echo.Context) error {
	var messages Message

	// Привязка входящих данных JSON к структуре Message
	if err := e.Bind(&messages); err != nil {
		return e.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Некорректный формат JSON",
		})
	}

	newTable := db.Create(&messages)
	if newTable != nil {
		fmt.Println("Меня так уже заебало обрабатывать эти ошибки ну когда это кончиться")
	} else {
		return e.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "GG я ливаю",
		})
	}

	return e.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Запись успешно добавлена в базу данных",
	})
}

var messages = make(map[int]Message)

func PutchHendler(e echo.Context) error {
	idParam := e.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "BAD PIVO",
		})
	}

	var upMessage Message

	if err := e.Bind(&upMessage); err != nil {
		return e.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Ну чё то ",
		})
	}

	var message Message
	result := db.First(&message, id)
	// if result.Error != gorm.ErrRecordNotFound {
	// 	return e.JSON(http.StatusNotFound, Response{
	// 		Status:  "Error",
	// 		Message: "Нет сообщения",
	// 	})
	// }

	if upMessage.Task != "" {
		message.Task = upMessage.Task
	}
	message.IsDone = upMessage.IsDone

	result = db.Save(&message)

	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Ошибка при обновлении записи",
		})
	}

	return e.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Запись успешно обновлена",
	})

}

func DeletHendler(e echo.Context) error {
	idParam := e.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, Response{
			Status:  "Error",
			Message: "Некорректный ID",
		})
	}

	var message Message
	result := db.First(&message, id)
	if result.Error != nil {
		return e.JSON(http.StatusNotFound, Response{
			Status:  "Error",
			Message: "Сообщение не найдено",
		})
	}

	result = db.Delete(&message)
	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, Response{
			Status:  "Error",
			Message: "Ошибка при удалении записи",
		})
	}

	return e.JSON(http.StatusOK, Response{
		Status:  "Success",
		Message: "Запись успешно удалена",
	})
}

func main() {
	InitDB()
	e := echo.New()

	e.GET("/messages", GetHandler)
	e.POST("/messages", PostHandler)
	e.PATCH("/messages/:id", PutchHendler)
	e.DELETE("/messages/:id", DeletHendler)
	e.Start("localhost:8080")

}

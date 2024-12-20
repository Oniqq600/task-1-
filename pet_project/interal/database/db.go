package database

import (
	"log"
	"pet_project/internal/taskServise"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {

	var err error
	dsn := "host=localhost user=postgres password=yourpassword dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Я не знаю что за фатал!!!", err)
	}

	err = DB.AutoMigrate(&orm.Message{})
	if err != nil {
		log.Fatal("Ошибка при миграции: ", err)
	}
}

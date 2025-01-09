package database

import (
	// orm "CRUD/internal/service/tasksService"
	user "CRUD/internal/service/usersService"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := "host=localhost user=postgres password=yourpassword dbname=main port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Fatal error", err)
	}

	err = DB.AutoMigrate(&user.Users{}, &user.Tasks{})
	if err != nil {
		panic("Migration error" + err.Error())
	}
}

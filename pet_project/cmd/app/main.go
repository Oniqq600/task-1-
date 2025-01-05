package main

import (
	"pet_project/internal/database"
	handlers "pet_project/internal/hendlers"
	taskServise "pet_project/internal/taskServise"
	orm "pet_project/internal/userService"

	"github.com/labstack/echo/v4/middleware"

	"pet_project/internal/web/tasks"
	user "pet_project/internal/web/users"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&taskServise.Message{}); err != nil {
		panic("Error")
	}

	repo := taskServise.NewTaskRepository(database.DB)
	service := taskServise.NewTaskService(repo)

	urepo := orm.NewUserRepository(database.DB)
	uservice := orm.NewUserService(urepo)

	handler := handlers.NewTaskHendler(service)
	uhandler := handlers.NewUserHandler(uservice)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	userHandler := user.NewStrictHandler(uhandler, nil)
	user.RegisterHandlers(e, userHandler)

	if err := e.Start("localhost:8080"); err != nil {
		panic("Error")
	}
}

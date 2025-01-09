package main

import (
	"CRUD/internal/database"
	handlers "CRUD/internal/hendlers"
	task_orm "CRUD/internal/service/tasksService"
	orm "CRUD/internal/service/usersService"

	"github.com/labstack/echo/v4/middleware"

	"CRUD/internal/web/tasks"
	user "CRUD/internal/web/users"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&orm.Tasks{}); err != nil {
		panic("Database connection error: " + err.Error())
	}

	repo := task_orm.NewTaskRepository(database.DB)
	service := task_orm.NewTaskService(repo)

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

	if err := e.Start("localhost:3000"); err != nil {
		panic("Error: " + err.Error())
	}

}

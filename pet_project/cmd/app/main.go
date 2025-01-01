package main

import (
	"pet_project/internal/database"
	"pet_project/internal/hendlers"
	orm "pet_project/internal/taskServise"

	"github.com/labstack/echo/v4/middleware"

	"pet_project/internal/web/tasks"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	if err := database.DB.AutoMigrate(&orm.Message{}); err != nil {
		panic("Error")
	}

	repo := orm.NewTaskRepository(database.DB)
	service := orm.NewService(repo)

	handler := hendlers.NewHendler(service)

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	strictHandler := tasks.NewStrictHandler(handler, nil)
	tasks.RegisterHandlers(e, strictHandler)

	e.GET("/api/get", handler.GetTaskHendler)
	e.POST("/api/post", handler.PostTaskHandler)
	e.PATCH("/api/patch/:id", handler.PutchTaskHendler)
	e.DELETE("/api/delete/:id", handler.DeleteTaskHendler)

	if err := e.Start("localhost:8080"); err != nil {
		panic("Error")
	}
}

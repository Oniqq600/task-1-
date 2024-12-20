package main

import (
	"pet_project/internal/database"
	"pet_project/internal/hendlers"
	orm "pet_project/internal/taskServise"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()
	repo := orm.NewTaskRepository(database.DB)
	service := orm.NewService(repo)

	hendler := hendlers.NewHendler(service)

	e := echo.New()

	e.GET("/api/get", hendler.GetTaskHendler)
	e.POST("/api/post", hendler.PostTaskHandler)
	e.PATCH("/api/patch/:id", hendler.PutchTaskHendler)
	e.DELETE("/api/delete/:id", hendler.DeleteTaskHendler)
	e.Start("localhost:8080")

}

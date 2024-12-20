package hendlers

import (
	"net/http"
	orm "pet_project/internal/taskServise"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service *orm.TaskService
}

func NewHendler(service *orm.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) GetTaskHendler(e echo.Context) error {
	task, err := h.Service.GetAllTasks()
	if err != nil {
		return e.String(http.StatusInternalServerError, err.Error())
	}
	return e.JSON(http.StatusOK, task)
}

func (h *Handler) PostTaskHandler(e echo.Context) error {
	var task orm.Message

	if err := e.Bind(&task); err != nil {
		return e.JSON(http.StatusBadRequest, orm.Response{
			Status:  "Error",
			Message: "Некорректный формат JSON",
		})
	}
	createdTask, err := h.Service.CreateTask(task)
	if err != nil {
		return e.JSON(http.StatusInternalServerError, orm.Response{
			Status:  "Фальшивка епта",
			Message: "Пару слов про это всё...",
		})
	}

	return e.JSON(http.StatusOK, createdTask)
}

func (h *Handler) PutchTaskHendler(e echo.Context) error {
	idParam := e.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, orm.Response{
			Status:  "Бабка: Это что у тебя. Оператор: BadRequest",
			Message: "Аааааа,  я думала сова",
		})
	}

	var task orm.Message
	if err := e.Bind(&task); err != nil {
		return e.JSON(http.StatusBadRequest, orm.Response{
			Status:  "Сегодня порция старых мемов",
			Message: "Шутка про повара",
		})
	}

	upTask, err := h.Service.UpdateTaskByID(uint(id), task)
	if err != nil {
		return e.JSON(http.StatusBadRequest, orm.Response{
			Status:  "Вооьще не за что я сидеть буду в этой тюрьме этой ",
			Message: "Шоколад ни в чём не виноват",
		})
	}

	return e.JSON(http.StatusOK, upTask)
}

func (h *Handler) DeleteTaskHendler(e echo.Context) error {
	idParam := e.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return e.JSON(http.StatusBadRequest, orm.Response{
			Status:  "Чисти",
			Message: "Как я буду вилкой то чистить",
		})
	}

	err = h.Service.DeleteTaskByID(uint(id))
	if err != nil {
		return e.JSON(http.StatusBadRequest, orm.Response{
			Status:  "Я устал мемы вспоминать",
			Message: "наверное не нашёл id нужный",
		})
	}
	return e.JSON(http.StatusOK, orm.Response{
		Status:  "It`s all good men",
		Message: "Better call Saul",
	})
}

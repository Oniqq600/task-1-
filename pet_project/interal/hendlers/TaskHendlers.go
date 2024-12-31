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

func (h *Handler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.Service.GetAllTasks()
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasks200JSONResponse{}

	for _, tsk := range allTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := orm.Message{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	response := tasks.PostTasks201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	taskRequest := request.Body

	taskToUpdate := orm.Message{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
	}

	updatedTask, err := h.Service.UpdateTaskByID(uint(taskID), taskToUpdate)
	if err != nil {
		return nil, err
	}

	response := tasks.PatchTasksId200JSONResponse{
		Id:     &updatedTask.ID,
		Task:   &updatedTask.Task,
		IsDone: &updatedTask.IsDone,
	}

	return response, nil
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	err := h.Service.DeleteTaskByID(uint(taskID))
	if err != nil {
		return nil, err
	}

	return &tasks.DeleteTasksId204Response{}, nil
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

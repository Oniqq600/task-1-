package handlers

import (
	orm "CRUD/internal/service/tasksService"
	user "CRUD/internal/service/usersService"
	"CRUD/internal/web/tasks"
	"context"
)

type Handler struct {
	Service *orm.TaskService
}

func NewTaskHendler(service *orm.TaskService) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	taskID := request.Id

	err := h.Service.DeleteTaskByID(uint(taskID))
	if err != nil {
		return nil, err
	}

	return &tasks.DeleteTasksId204Response{}, nil
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

func (h *Handler) GetTasksUserUserId(ctx context.Context, request tasks.GetTasksUserUserIdRequestObject) (tasks.GetTasksUserUserIdResponseObject, error) {
	userID := request.UserId
	userTasks, err := h.Service.GetTasksByUserID(uint(userID))
	if err != nil {
		return nil, err
	}

	response := tasks.GetTasksUserUserId200JSONResponse{}
	for _, tsk := range userTasks {
		task := tasks.Task{
			Id:     &tsk.ID,
			Task:   &tsk.Task,
			IsDone: &tsk.IsDone,
		}
		response = append(response, task)
	}

	return response, nil
}

func (h *Handler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	taskID := request.Id
	taskRequest := request.Body

	taskToUpdate := user.Tasks{
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

func (h *Handler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := user.Tasks{
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

func (h *Handler) PostTasksUserId(ctx context.Context, request tasks.PostTasksUserIdRequestObject) (tasks.PostTasksUserIdResponseObject, error) {
	userID := request.UserId
	taskRequest := request.Body

	taskToCreate := user.Tasks{
		Task:   *taskRequest.Task,
		IsDone: *taskRequest.IsDone,
		UserID: userID,
	}
	createdTask, err := h.Service.CreateTask(taskToCreate)
	if err != nil {
		return nil, err
	}

	// Формируем ответ
	response := tasks.PostTasksUserId201JSONResponse{
		Id:     &createdTask.ID,
		Task:   &createdTask.Task,
		IsDone: &createdTask.IsDone,
	}

	return response, nil
}

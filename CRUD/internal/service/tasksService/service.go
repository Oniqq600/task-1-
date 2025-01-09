package orm

import (
	orm "CRUD/internal/service/usersService"
)

type TaskService struct {
	repo MessageRepository
}

func NewTaskService(repo *taskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task orm.Tasks) (orm.Tasks, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]orm.Tasks, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, task orm.Tasks) (orm.Tasks, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

func (s *TaskService) GetTasksByUserID(userID uint) ([]orm.Tasks, error) {
	return s.repo.GetTasksByUserID(userID)
}

func (s *TaskService) CreateTaskForUser(userID uint, task orm.Tasks) (orm.Tasks, error) {
	task.UserID = int(userID)
	return s.repo.CreateTask(task)
}

package taskServise

type TaskService struct {
	repo MessageRepository
}

func NewTaskService(repo *taskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Message) (Message, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Message, error) {
	return s.repo.GetAllTasks()
}

func (s *TaskService) UpdateTaskByID(id uint, task Message) (Message, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskByID(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

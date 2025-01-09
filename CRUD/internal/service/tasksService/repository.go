package orm

import (
	orm "CRUD/internal/service/usersService"

	"gorm.io/gorm"
)

type MessageRepository interface {
	GetAllTasks() ([]orm.Tasks, error)
	CreateTask(task orm.Tasks) (orm.Tasks, error)
	UpdateTaskByID(id uint, task orm.Tasks) (orm.Tasks, error)
	DeleteTaskByID(id uint) error
	GetTasksByUserID(userID uint) ([]orm.Tasks, error)
	PostTaskWithUser(task orm.Tasks) (orm.Tasks, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task orm.Tasks) (orm.Tasks, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return orm.Tasks{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]orm.Tasks, error) {
	var tasks []orm.Tasks
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task orm.Tasks) (orm.Tasks, error) {
	var existingTask orm.Tasks
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return orm.Tasks{}, err
	}

	if err := r.db.Model(&existingTask).Updates(task).Error; err != nil {
		return orm.Tasks{}, err
	}

	return existingTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	var existingTask orm.Tasks
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&existingTask).Error; err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) GetTasksByUserID(userID uint) ([]orm.Tasks, error) {
	var tasks []orm.Tasks
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) PostTaskWithUser(task orm.Tasks) (orm.Tasks, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return orm.Tasks{}, result.Error
	}
	return task, nil
}

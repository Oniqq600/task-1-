package taskServise

import (
	"gorm.io/gorm"
)

type MessageRepository interface {
	GetAllTasks() ([]Message, error)

	CreateTask(task Message) (Message, error)

	UpdateTaskByID(id uint, task Message) (Message, error)

	DeleteTaskByID(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *taskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) CreateTask(task Message) (Message, error) {
	result := r.db.Create(&task)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return task, nil
}

func (r *taskRepository) GetAllTasks() ([]Message, error) {
	var tasks []Message
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) UpdateTaskByID(id uint, task Message) (Message, error) {
	var existingTask Message
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return Message{}, err
	}

	if err := r.db.Model(&existingTask).Updates(task).Error; err != nil {
		return Message{}, err
	}

	return existingTask, nil
}

func (r *taskRepository) DeleteTaskByID(id uint) error {
	var existingTask Message
	if err := r.db.First(&existingTask, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&existingTask).Error; err != nil {
		return err
	}

	return nil
}

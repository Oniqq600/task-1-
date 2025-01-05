package orm

import (
	"gorm.io/gorm"
)

type UsersRepository interface {
	GetUsers() ([]Users, error)

	PostUser(user Users) (Users, error)

	PatchUserByID(id uint, user Users) (Users, error)

	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUsers() ([]Users, error) {
	var user []Users
	err := r.db.Find(&user).Error
	return user, err
}

func (r *userRepository) PostUser(user Users) (Users, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return Users{}, result.Error
	}
	return user, nil
}

func (r *userRepository) PatchUserByID(id uint, user Users) (Users, error) {
	var existingUser Users
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return Users{}, err
	}
	if err := r.db.Model(&existingUser).Updates(user).Error; err != nil {
		return Users{}, err
	}

	return existingUser, nil
}

func (r *userRepository) DeleteUserByID(id uint) error {
	var existingUser Users
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return err
	}

	if err := r.db.Delete(&existingUser).Error; err != nil {
		return err
	}

	return nil
}

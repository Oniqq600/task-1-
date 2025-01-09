package orm

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	// Указание на связь "один ко многим" с таблицей tasks
	Tasks []Tasks `gorm:"foreignKey:UserID;references:ID" json:"tasks"`
}

type Tasks struct {
	gorm.Model
	Task   string `json:"task"`
	IsDone bool   `json:"is_done"`
	UserID int    `gorm:"index" json:"user_id"`
	// Указание на связь с пользователем
	User Users `gorm:"foreignKey:UserID" json:"user"`
}

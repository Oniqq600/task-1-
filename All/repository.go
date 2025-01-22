package legal_entities

import (
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GetAll(ctx context.Context) ([]LegalEntities, error)
	Create(ctx context.Context, entity LegalEntities) error
	Update(ctx context.Context, entity LegalEntities) error
	Delete(ctx context.Context, uuid string) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]LegalEntities, error) {
	var entities []LegalEntities
	if err := r.db.WithContext(ctx).Find(&entities).Error; err != nil {
		return nil, err
	}
	return entities, nil
}

func (r *repository) Create(ctx context.Context, entity LegalEntities) error {
	return r.db.WithContext(ctx).Create(&entity).Error
}

func (r *repository) Update(ctx context.Context, entity LegalEntities) error {
	return r.db.WithContext(ctx).Save(&entity).Error
}

func (r *repository) Delete(ctx context.Context, uuid string) error {
	return r.db.WithContext(ctx).Where("uuid = ?", uuid).Delete(&LegalEntities{}).Error
}

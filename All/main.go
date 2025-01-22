package legal_entities

import "context"

type Service interface {
	GetAllLegalEntities(ctx context.Context) ([]LegalEntities, error)
	CreateLegalEntities(ctx context.Context, entity LegalEntities) error
	UpdateLegalEntities(ctx context.Context, entity LegalEntities) error
	DeleteLegalEntities(ctx context.Context, uuid string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetAllLegalEntities(ctx context.Context) ([]LegalEntities, error) {
	return s.repo.GetAll(ctx)
}

func (s *service) CreateLegalEntities(ctx context.Context, entity LegalEntities) error {
	return s.repo.Create(ctx, entity)
}

func (s *service) UpdateLegalEntities(ctx context.Context, entity LegalEntities) error {
	return s.repo.Update(ctx, entity)
}

func (s *service) DeleteLegalEntities(ctx context.Context, uuid string) error {
	return s.repo.Delete(ctx, uuid)
}

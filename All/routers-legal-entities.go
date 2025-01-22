package web

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/krisch/crm-backend/internal/jwt"
	"github.com/krisch/crm-backend/internal/legal_entities"
	oapi "github.com/krisch/crm-backend/internal/web/ofederation"
	echo "github.com/labstack/echo/v4"

	// "github.com/samber/lo".
	"github.com/sirupsen/logrus"
)

func initOpenAPILegaEntityRouters(a *Web, e *echo.Echo) {
	logrus.WithField("route", "oProject").Debug("routes initialization")

	midlewares := []oapi.StrictMiddlewareFunc{
		ValidateStructMiddeware,
		AuthMiddeware(a.app, []string{}),
	}

	handlers := oapi.NewStrictHandler(a, midlewares)
	oapi.RegisterHandlers(e, handlers)
}

func (a *Web) DeleteLegalEntitiesUuid(ctx context.Context, request oapi.DeleteLegalEntitiesUuidRequestObject) (oapi.DeleteLegalEntitiesUuidResponseObject, error) {
	_, ok := ctx.Value(claimsKey).(jwt.Claims)
	if !ok {
		return nil, ErrInvalidAuthHeader
	}

	err := a.app.LegalEntitiesService.DeleteLegalEntities(ctx, request.Uuid)
	if err != nil {
		return nil, err
	}

	return oapi.DeleteLegalEntitiesUuid204Response{}, nil
}

func toStringPointer(s string) *string {
	return &s
}

func (a *Web) GetLegalEntitiesUuid(ctx context.Context, request oapi.GetLegalEntitiesRequestObject) (oapi.GetLegalEntitiesResponseObject, error) {
	_, ok := ctx.Value(claimsKey).(jwt.Claims)
	if !ok {
		return nil, ErrInvalidAuthHeader
	}

	entities, err := a.app.LegalEntitiesService.GetAllLegalEntities(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve entity: %w", err)
	}

	items := make([]oapi.LegalEntityDTO, len(entities))
	for i, entity := range entities {
		items[i] = oapi.LegalEntityDTO{
			Uuid:      toStringPointer(entity.UUID.String()),
			Name:      &entity.Name,
			CreatedAt: &entity.CreatedAt,
			UpdatedAt: &entity.UpdatedAt,
			DeletedAt: &entity.DeletedAt,
		}
	}

	return oapi.GetLegalEntities200JSONResponse{
		Items: &items,
	}, nil
}

func (a *Web) PostLegalEntities(ctx context.Context, request oapi.PostLegalEntitiesRequestObject) (oapi.PostLegalEntitiesResponseObject, error) {
	_, ok := ctx.Value(claimsKey).(jwt.Claims)
	if !ok {
		return nil, ErrInvalidAuthHeader
	}

	if request.Body.Uuid == nil {
		return nil, fmt.Errorf("UUID is required")
	}

	uuidValue, err := uuid.Parse(*request.Body.Uuid)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	if request.Body.Name == nil {
		return nil, fmt.Errorf("Name is required")
	}

	entity := legal_entities.LegalEntities{
		Name: *request.Body.Name,
	}
	if err := a.app.LegalEntitiesService.CreateLegalEntities(ctx, entity); err != nil {
		return nil, fmt.Errorf("failed to create legal entity: %w", err)
	}

	uuidStr := uuidValue.String()

	return oapi.PostLegalEntities201JSONResponse{
		Uuid: &uuidStr,
	}, nil
}

func (a *Web) PutLegalEntitiesUuid(ctx context.Context, request oapi.PutLegalEntitiesUuidRequestObject) (oapi.PutLegalEntitiesUuidResponseObject, error) {
	_, ok := ctx.Value(claimsKey).(jwt.Claims)
	if !ok {
		return nil, ErrInvalidAuthHeader
	}

	uuidValue, err := uuid.Parse(request.Uuid)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID", err)
	}

	entity := legal_entities.LegalEntities{
		UUID: uuidValue,
		Name: *request.Body.Name,
	}
	if err := a.app.LegalEntitiesService.UpdateLegalEntities(ctx, entity); err != nil {
		return nil, err
	}

	return oapi.PutLegalEntitiesUuid200Response{}, nil
}

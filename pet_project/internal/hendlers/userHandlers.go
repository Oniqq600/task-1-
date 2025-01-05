package handlers

import (
	"context"
	orm "pet_project/internal/userService"
	user "pet_project/internal/web/users"
)

type UserHandler struct {
	Service *orm.UserService
}

func NewUserHandler(service *orm.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func (h *UserHandler) GetUsers(ctx context.Context, request user.GetUsersRequestObject) (user.GetUsersResponseObject, error) {
	allUsers, err := h.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	response := user.GetUsers200JSONResponse{}

	for _, usr := range allUsers {
		id := int(usr.ID)
		user := user.User{
			Id:       &id,
			Email:    &usr.Email,
			Password: &usr.Password,
		}
		response = append(response, user)
	}

	return response, nil
}

func (h *UserHandler) PostUsers(ctx context.Context, request user.PostUsersRequestObject) (user.PostUsersResponseObject, error) {
	userRequest := request.Body

	userToCreate := orm.Users{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	createdUser, err := h.Service.CreateUser(userToCreate)
	if err != nil {
		return nil, err
	}

	id := int(createdUser.ID)
	response := user.PostUsers201JSONResponse{
		Id:       &id,
		Email:    &createdUser.Email,
		Password: &createdUser.Password,
	}

	return response, nil
}

func (h *UserHandler) PatchUsersId(ctx context.Context, request user.PatchUsersIdRequestObject) (user.PatchUsersIdResponseObject, error) {
	userID := request.Id
	userRequest := request.Body

	userToUpdate := orm.Users{
		Email:    *userRequest.Email,
		Password: *userRequest.Password,
	}

	updatedUser, err := h.Service.UpdateUserByID(uint(userID), userToUpdate)
	if err != nil {
		return nil, err
	}

	id := int(updatedUser.ID)
	response := user.PatchUsersId200JSONResponse{
		Id:       &id,
		Email:    &updatedUser.Email,
		Password: &updatedUser.Password,
	}

	return response, nil
}

func (h *UserHandler) DeleteUsersId(ctx context.Context, request user.DeleteUsersIdRequestObject) (user.DeleteUsersIdResponseObject, error) {
	userID := request.Id

	err := h.Service.DeleteUserByID(uint(userID))
	if err != nil {
		return nil, err
	}

	return &user.DeleteUsersId204Response{}, nil
}

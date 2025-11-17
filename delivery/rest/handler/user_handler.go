package handler

import (
	"fmt"
	"net/http"
	"user-management/internal/service"
	"user-management/pkg/dto"
	"user-management/pkg/response"

	"github.com/goccy/go-json"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
func (h *UserHandler) UserChangeProfile(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	request := new(dto.UserUpdateRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		fmt.Println(err.Error())
		panic(response.Except(http.StatusBadRequest, "failed decode body to json"))
	}
	err := h.userService.UserChangeProfile(id, request)
	if err != nil {
		panic(err)
	}
	response.Success(w, 200, "OK", r.URL.Path)
}
func (h *UserHandler) UserGetByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	result, err := h.userService.UserGetByUsername(username)
	if err != nil {
		panic(err)
	}
	response.Success(w, 200, result, r.URL.Path)
}
func (h *UserHandler) UserGetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	result, err := h.userService.UserGetById(id)
	if err != nil {
		panic(err)
	}
	response.Success(w, 200, result, r.URL.Path)
}

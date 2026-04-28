package handler

import (
	"encoding/json"
	"fmt"
	"go-polls-service/internal/core/core_logger"
	"go-polls-service/internal/core/dto"
	"go-polls-service/internal/core/http_server"
	"net/http"
)

type UsersHTTPHandler struct {
	usersService UsersService
}

type UsersService interface{}

func NewUsersHTTPHandler(usersService UsersService) *UsersHTTPHandler {
	return &UsersHTTPHandler{
		usersService: usersService,
	}
}

func (h *UsersHTTPHandler) Routes() []http_server.Route {
	return []http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/users",
			Handler: h.CreateUser,
		},
	}
}

func (h *UsersHTTPHandler) CreateUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)

	log.Debug("Invoke CreateUser Handler")

	var request dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		fmt.Println("error:", err)
	}

	rw.WriteHeader(http.StatusOK)
}

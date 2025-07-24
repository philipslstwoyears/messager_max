package apiServer

import (
	"awesomeProject2/service"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

type RegisterRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

func (handler *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logrus.Error()
		return
	}
	err = handler.service.RegisterUser(req.Login, req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		logrus.Error()
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user registered successfully"))
}

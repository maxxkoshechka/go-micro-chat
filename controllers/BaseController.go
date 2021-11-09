package controllers

import (
	"fmt"
	"net/http"

	"chat/models"
)

// BaseHandler will hold everything that controller needs
type BaseHandler struct {
	UserRepo models.UserRepository
}

// NewBaseHandler returns a new BaseHandler
func NewBaseHandler(userRepo models.UserRepository) *BaseHandler {
	return &BaseHandler{
		UserRepo: userRepo,
	}
}

// HelloWorld returns Hello, World
func (h *BaseHandler) HelloWorld(w http.ResponseWriter, r *http.Request) {
	if user, err := h.UserRepo.Select("1"); err != nil {
		fmt.Println("Error", user)
	}

	w.Write([]byte("Hello, World"))
}

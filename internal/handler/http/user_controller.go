package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/wonjinsin/go-boilerplate/internal/handler/http/dto"
	"github.com/wonjinsin/go-boilerplate/internal/usecase"
	"github.com/wonjinsin/go-boilerplate/pkg/utils"
)

// UserController handles user-related HTTP requests
type UserController struct {
	svc usecase.UserService
}

// NewUserController creates a new user controller
func NewUserController(svc usecase.UserService) *UserController {
	return &UserController{svc: svc}
}

// CreateUser handles user creation
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.WriteStandardJSON(w, r, http.StatusBadRequest, dto.ErrorResult{
			Msg: "invalid json",
		})
		return
	}

	u, err := c.svc.CreateUser(r.Context(), req.Name, req.Email)
	if err != nil {
		writeError(w, r, err)
		return
	}

	response := dto.ToUserResponse(u)
	utils.WriteStandardJSON(w, r, http.StatusCreated, response)
}

// ListUsers handles user listing with pagination
func (c *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	offset, limit := utils.ParsePagination(r)
	list, err := c.svc.ListUsers(r.Context(), offset, limit)
	if err != nil {
		utils.WriteStandardJSON(w, r, http.StatusInternalServerError, dto.ErrorResult{
			Msg: err.Error(),
		})
		return
	}

	response := dto.ToUserListResponse(list, len(list), offset, limit)
	utils.WriteStandardJSON(w, r, http.StatusOK, response)
}

// GetUser handles retrieving a single user by ID
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.WriteStandardJSON(w, r, http.StatusBadRequest, dto.ErrorResult{
			Msg: "user id is required",
		})
		return
	}

	u, err := c.svc.GetUser(r.Context(), id)
	if err != nil {
		writeError(w, r, err)
		return
	}

	response := dto.ToUserResponse(u)
	utils.WriteStandardJSON(w, r, http.StatusOK, response)
}

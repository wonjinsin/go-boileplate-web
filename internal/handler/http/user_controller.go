package http

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/wonjinsin/go-boilerplate/internal/constants"
	"github.com/wonjinsin/go-boilerplate/internal/handler/http/dto"
	"github.com/wonjinsin/go-boilerplate/internal/interfaces"
	"github.com/wonjinsin/go-boilerplate/pkg/utils"
)

// UserController handles user-related HTTP requests
type UserController struct {
	svc interfaces.UserService
}

// NewUserController creates a new user controller
func NewUserController(svc interfaces.UserService) *UserController {
	return &UserController{svc: svc}
}

// CreateUser handles user creation
func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := utils.ParseJSONBody(r, &req); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Error: "invalid json",
			Code:  constants.ErrCodeInvalidJSON,
		})
		return
	}

	u, err := c.svc.CreateUser(req.Name, req.Email)
	if err != nil {
		writeError(w, err)
		return
	}

	response := dto.ToUserResponse(u)
	utils.WriteJSON(w, http.StatusCreated, response)
}

// ListUsers handles user listing with pagination
func (c *UserController) ListUsers(w http.ResponseWriter, r *http.Request) {
	offset, limit := utils.ParsePagination(r)
	list, err := c.svc.ListUsers(offset, limit)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, dto.ErrorResponse{
			Error: err.Error(),
			Code:  constants.ErrCodeInternalError,
		})
		return
	}

	response := dto.ToUserListResponse(list, len(list), offset, limit)
	utils.WriteJSON(w, http.StatusOK, response)
}

// GetUser handles retrieving a single user by ID
func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Error: "user id is required",
			Code:  constants.ErrCodeMissingUserID,
		})
		return
	}

	u, err := c.svc.GetUser(id)
	if err != nil {
		writeError(w, err)
		return
	}

	response := dto.ToUserResponse(u)
	utils.WriteJSON(w, http.StatusOK, response)
}

package httpapi

import (
	"net/http"

	"example/internal/adapter/httpapi/model"
	"example/internal/constants"
	"example/internal/domain"
	"example/internal/usecase"
	"example/pkg/utils"
)

type userHandler struct {
	svc usecase.UserService
}

func RegisterUserRoutes(mux *http.ServeMux, svc usecase.UserService) {
	h := &userHandler{svc: svc}

	// Health check
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// User endpoints
	mux.HandleFunc("/users", h.users)
	mux.HandleFunc("/users/", h.userByID)
}

func (h *userHandler) users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req model.CreateUserRequest
		if err := utils.ParseJSONBody(r, &req); err != nil {
			utils.WriteJSON(w, http.StatusBadRequest, model.ErrorResponse{
				Error: "invalid json",
				Code:  constants.ErrCodeInvalidJSON,
			})
			return
		}

		u, err := h.svc.CreateUser(req.Name, req.Email)
		if err != nil {
			writeDomainError(w, err)
			return
		}

		response := model.ToUserResponse(u)
		utils.WriteJSON(w, http.StatusCreated, response)

	case http.MethodGet:
		offset, limit := utils.ParsePagination(r)
		list, err := h.svc.ListUsers(offset, limit)
		if err != nil {
			utils.WriteJSON(w, http.StatusInternalServerError, model.ErrorResponse{
				Error: err.Error(),
				Code:  constants.ErrCodeInternalError,
			})
			return
		}

		response := model.ToUserListResponse(list, len(list), offset, limit)
		utils.WriteJSON(w, http.StatusOK, response)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *userHandler) userByID(w http.ResponseWriter, r *http.Request) {
	id := utils.ExtractPathParam(r.URL.Path, "/users/")
	if id == "" {
		utils.WriteJSON(w, http.StatusNotFound, model.ErrorResponse{
			Error: "user id is required",
			Code:  constants.ErrCodeMissingUserID,
		})
		return
	}

	switch r.Method {
	case http.MethodGet:
		u, err := h.svc.GetUser(id)
		if err != nil {
			writeDomainError(w, err)
			return
		}

		response := model.ToUserResponse(u)
		utils.WriteJSON(w, http.StatusOK, response)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func writeDomainError(w http.ResponseWriter, err error) {
	switch err {
	case domain.ErrUserNotFound:
		utils.WriteJSON(w, http.StatusNotFound, model.ErrorResponse{
			Error: err.Error(),
			Code:  constants.ErrCodeUserNotFound,
		})
	case domain.ErrDuplicateEmail:
		utils.WriteJSON(w, http.StatusConflict, model.ErrorResponse{
			Error: err.Error(),
			Code:  constants.ErrCodeDuplicateEmail,
		})
	case domain.ErrInvalidName:
		utils.WriteJSON(w, http.StatusBadRequest, model.ErrorResponse{
			Error: err.Error(),
			Code:  constants.ErrCodeInvalidName,
		})
	default:
		utils.WriteJSON(w, http.StatusBadRequest, model.ErrorResponse{
			Error: err.Error(),
			Code:  constants.ErrCodeBadRequest,
		})
	}
}

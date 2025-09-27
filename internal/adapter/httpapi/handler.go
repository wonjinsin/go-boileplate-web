package httpapi

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"example/internal/adapter/httpapi/model"
	"example/internal/domain"
	"example/internal/usecase"
)

type userHandler struct {
	svc usecase.UserService
}

func RegisterUserRoutes(mux *http.ServeMux, svc usecase.UserService) {
	h := &userHandler{svc: svc}
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })
	mux.HandleFunc("/users", h.users)
	mux.HandleFunc("/users/", h.userByID)
}

func (h *userHandler) users(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var req model.CreateUserRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			writeJSON(w, http.StatusBadRequest, model.ErrorResponse{
				Error: "invalid json",
				Code:  "INVALID_JSON",
			})
			return
		}

		u, err := h.svc.CreateUser(req.Name, req.Email)
		if err != nil {
			writeDomainError(w, err)
			return
		}

		response := model.ToUserResponse(u)
		writeJSON(w, http.StatusCreated, response)

	case http.MethodGet:
		offset, limit := parsePagination(r)
		list, err := h.svc.ListUsers(offset, limit)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, model.ErrorResponse{
				Error: err.Error(),
				Code:  "INTERNAL_ERROR",
			})
			return
		}

		response := model.ToUserListResponse(list, len(list), offset, limit)
		writeJSON(w, http.StatusOK, response)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *userHandler) userByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/users/"):]
	if id == "" {
		writeJSON(w, http.StatusNotFound, model.ErrorResponse{
			Error: "user id is required",
			Code:  "MISSING_USER_ID",
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
		writeJSON(w, http.StatusOK, response)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func parsePagination(r *http.Request) (int, int) {
	q := r.URL.Query()
	offset, _ := strconv.Atoi(q.Get("offset"))
	limit, _ := strconv.Atoi(q.Get("limit"))
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 || limit > 200 {
		limit = 50
	}
	return offset, limit
}

func writeJSON(w http.ResponseWriter, code int, v any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Printf("json encode error: %v", err)
	}
}

func writeDomainError(w http.ResponseWriter, err error) {
	switch err {
	case domain.ErrUserNotFound:
		writeJSON(w, http.StatusNotFound, model.ErrorResponse{
			Error: err.Error(),
			Code:  "USER_NOT_FOUND",
		})
	case domain.ErrDuplicateEmail:
		writeJSON(w, http.StatusConflict, model.ErrorResponse{
			Error: err.Error(),
			Code:  "DUPLICATE_EMAIL",
		})
	case domain.ErrInvalidName:
		writeJSON(w, http.StatusBadRequest, model.ErrorResponse{
			Error: err.Error(),
			Code:  "INVALID_NAME",
		})
	default:
		writeJSON(w, http.StatusBadRequest, model.ErrorResponse{
			Error: err.Error(),
			Code:  "BAD_REQUEST",
		})
	}
}

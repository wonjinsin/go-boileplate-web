package http

import (
	"net/http"

	"example/internal/constants"
	"example/internal/domain"
	"example/internal/handler/http/dto"
	"example/pkg/utils"
)

// writeError writes domain errors as HTTP responses
func writeError(w http.ResponseWriter, err error) {
	switch err {
	case domain.ErrUserNotFound:
		utils.WriteJSON(w, http.StatusNotFound, dto.ErrorResponse{
			Error: err.Error(),
			Code:  constants.ErrCodeUserNotFound,
		})
	case domain.ErrDuplicateEmail:
		utils.WriteJSON(w, http.StatusConflict, dto.ErrorResponse{
			Error: err.Error(),
			Code:  constants.ErrCodeDuplicateEmail,
		})
	case domain.ErrInvalidName:
		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
			Code:  constants.ErrCodeInvalidName,
		})
	default:
		utils.WriteJSON(w, http.StatusBadRequest, dto.ErrorResponse{
			Error: err.Error(),
			Code:  constants.ErrCodeBadRequest,
		})
	}
}


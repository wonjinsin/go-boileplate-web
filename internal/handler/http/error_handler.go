package http

import (
	"net/http"

	"github.com/wonjinsin/go-boilerplate/internal/domain"
	"github.com/wonjinsin/go-boilerplate/internal/handler/http/dto"
	"github.com/wonjinsin/go-boilerplate/pkg/utils"
)

// writeError writes domain errors as HTTP responses
func writeError(w http.ResponseWriter, r *http.Request, err error) {
	switch err {
	case domain.ErrUserNotFound:
		utils.WriteStandardJSON(w, r, http.StatusNotFound, dto.ErrorResult{
			Msg: err.Error(),
		})
	case domain.ErrDuplicateEmail:
		utils.WriteStandardJSON(w, r, http.StatusConflict, dto.ErrorResult{
			Msg: err.Error(),
		})
	case domain.ErrInvalidName:
		utils.WriteStandardJSON(w, r, http.StatusBadRequest, dto.ErrorResult{
			Msg: err.Error(),
		})
	default:
		utils.WriteStandardJSON(w, r, http.StatusBadRequest, dto.ErrorResult{
			Msg: err.Error(),
		})
	}
}

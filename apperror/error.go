package apperror

import "net/http"

type AppError struct {
	Id       string
	HttpCode int
}

func (e *AppError) Error() string {
	return e.Id
}

var (
	InternalError = &AppError{"Internal error", http.StatusInternalServerError}
	Unauthorized  = &AppError{"Unauthorized", http.StatusUnauthorized}
)

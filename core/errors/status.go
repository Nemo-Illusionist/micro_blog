package errors

import "net/http"

func NewBoom(code int, message string, err error) (int, Boom) {
	return code, Boom{code, message, err}
}

func BadRequestBoom(err error) (int, Boom) {
	return NewBoom(http.StatusBadRequest, "bad request", err)
}

func InternalServerErrorBoom(err error) (int, Boom) {
	return NewBoom(http.StatusInternalServerError, "internal server error", err)
}

func NotFoundBoom() (int, Boom) {
	return NewBoom(http.StatusNotFound, "object not found", nil)
}

func ForbiddenBoom() (int, Boom) {
	return NewBoom(http.StatusForbidden, "forbidden", nil)
}

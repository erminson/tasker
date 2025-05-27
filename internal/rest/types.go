package rest

import (
	"encoding/json"
	"io"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Status  int    `json:"status"`
	Data    any    `json:"data,omitempty"`
}

type HandlerFunc[T any, R any] func(r *http.Request, params T) (R, error)

type Validator interface {
	Valid() error
}

func WrapHandler[T any, R any](h HandlerFunc[T, R]) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params T
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			writeError(w, http.StatusBadRequest, err)
			return
		}

		data, err := h(r, params)
		if err != nil {
			writeError(w, http.StatusInternalServerError, err)
			return
		}

		OK(w, data)
	}
}

func BadRequest(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(Response{
		Error:  err,
		Status: http.StatusBadRequest,
	})
}

func NotFound(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	_ = json.NewEncoder(w).Encode(Response{
		Error:  err.Error(),
		Status: http.StatusNotFound,
	})
}

func WriteError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Response{
		Error:  err.Error(),
		Status: status,
	})
}

func Internal(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(Response{
		Error:  err.Error(),
		Status: http.StatusInternalServerError,
	})
}

func Unauthorized(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusUnauthorized)
	_ = json.NewEncoder(w).Encode(Response{
		Error:  err,
		Status: http.StatusUnauthorized,
	})
}

func Forbidden(w http.ResponseWriter, err string) {
	w.WriteHeader(http.StatusForbidden)
	_ = json.NewEncoder(w).Encode(Response{
		Error:  err,
		Status: http.StatusForbidden,
	})
}

func OK(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(Response{
		Status:  http.StatusOK,
		Data:    data,
		Success: true,
	})
}

func Accepted(w http.ResponseWriter, data any) {
	w.WriteHeader(http.StatusAccepted)
	_ = json.NewEncoder(w).Encode(Response{
		Status:  http.StatusAccepted,
		Data:    data,
		Success: true,
	})
}

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
	_ = json.NewEncoder(w).Encode(Response{
		Status:  http.StatusNoContent,
		Success: true,
	})
}

func writeError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(Response{
		Error:  err.Error(),
		Status: status,
	})
}

func DecodeJson[T any](body io.ReadCloser) (T, error) {
	var result T
	err := json.NewDecoder(body).Decode(&result)
	return result, err
}

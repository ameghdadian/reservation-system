package response

import (
	"encoding/json"
	"errors"

	"github.com/ameghdadian/service/business/data/page"
)

type PageDocument[T any] struct {
	Items       []T `json:"items"`
	Total       int `json:"total"`
	Page        int `json:"page"`
	RowsPerPage int `json:"rows_per_page"`
}

func NewPageDocument[T any](items []T, total int, page page.Page) PageDocument[T] {
	return PageDocument[T]{
		Items:       items,
		Total:       total,
		Page:        page.Number(),
		RowsPerPage: page.RowsPerPage(),
	}
}

func (pd PageDocument[T]) Encode() ([]byte, string, error) {
	data, err := json.Marshal(pd)
	return data, "application/json", err
}

// ============================================================================

// ErrorDocument is the form used for API responses from failures in the API.
type ErrorDocument struct {
	Error  string            `json:"error"`
	Fields map[string]string `json:"fields,omitempty"`
}

// Error is used to pass an error during the request through the
// application with web specific context.
type Error struct {
	Err    error
	Status int
}

func NewError(err error, status int) error {
	return &Error{err, status}
}

func (re *Error) Error() string {
	return re.Err.Error()
}

func IsError(err error) bool {
	var re *Error
	return errors.As(err, &re)
}

func GetError(err error) *Error {
	var re *Error
	if !errors.As(err, &re) {
		return nil
	}

	return re
}

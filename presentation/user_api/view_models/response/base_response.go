package response

import (
	"time"
)

type BaseResponse[T any] struct {
	Date    time.Time `json:"date"`
	Data    T         `json:"data"`
	Message string    `json:"message"`
}

type ErrorResponse struct {
	Date    time.Time `json:"date"`
	Message string    `json:"message"`
}

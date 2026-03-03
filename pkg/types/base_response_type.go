package types

import (
	"gear-server/internal/enum"
)

type Meta struct {
	Page       int   `json:"page,omitempty"`
	PageSize   int   `json:"page_size,omitempty"`
	Total      int64 `json:"total,omitempty"`
	TotalPages int   `json:"total_pages,omitempty"`
	HasNext    bool  `json:"has_next,omitempty"`
	HasPrev    bool  `json:"has_prev,omitempty"`
}

type BaseResponse[T any] struct {
	Code    enum.ResponseCode `json:"code"`
	Message string            `json:"message"`
	Data    T                 `json:"data,omitempty"`
	Meta    *Meta             `json:"meta,omitempty"`
	Error   interface{}       `json:"error,omitempty"`
}

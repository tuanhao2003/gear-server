package helper

import (
	"gear-server/internal/enum"
	"gear-server/pkg/types"
)

type ResponseOption[T any] func(*types.BaseResponse[T])

func WithMeta[T any](meta *types.Meta) ResponseOption[T] {
	return func(r *types.BaseResponse[T]) {
		r.Meta = meta
	}
}

func WithError[T any](err interface{}) ResponseOption[T] {
	return func(r *types.BaseResponse[T]) {
		r.Error = err
	}
}

func Response[T any](
	code enum.ResponseCode,
	message string,
	data T,
	opts ...ResponseOption[T],
) types.BaseResponse[T] {

	res := types.BaseResponse[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}

	for _, opt := range opts {
		opt(&res)
	}

	return res
}

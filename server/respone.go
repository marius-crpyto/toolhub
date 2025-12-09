package server

type Response[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func Ok[T any](data T) Response[T] {
	return Response[T]{Code: 0, Data: data}
}

func Err(code int, msg string) Response[struct{}] {
	return Response[struct{}]{Code: code, Message: msg}
}

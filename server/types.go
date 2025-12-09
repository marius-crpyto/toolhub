package server

import "github.com/gin-gonic/gin"

type ApiPath struct {
	Path    string          // path
	Method  string          // method
	Handler gin.HandlerFunc // handler func

	AuthMiddleware func(c *gin.Context) error // auth middleware
}

type FilterHandle func(c *gin.Context) error

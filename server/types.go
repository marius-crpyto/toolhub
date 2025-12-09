package server

import "github.com/gin-gonic/gin"

const (
	GET    string = "GET"
	POST   string = "POST"
	PUT    string = "PUT"
	DELETE string = "DELETE"
)

type ApiPath struct {
	Path    string          // path
	Method  string          // method
	Handler gin.HandlerFunc // handler func

	Auth func(c *gin.Context) error // auth middleware
}

type FilterHandler func(c *gin.Context) error

func WrapFilter(h FilterHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			c.AbortWithStatusJSON(401, Err(401, err.Error()))
			return
		}
		c.Next()
	}
}

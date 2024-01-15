package middleware

import (
	"github.com/gin-gonic/gin"
)

func GinErrorHandle(h func(c *gin.Context) error) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := h(c); err != nil {
			if len(c.Errors) == 0 {
				c.Error(err)
			}
		}
	}
}

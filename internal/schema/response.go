package schema

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Empty = struct{}

//swagger:model http_response
type Response[T any] struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Result  T      `json:"result"`
}

func Respond[T any](v T, c *gin.Context) error {
	c.JSON(http.StatusOK, Response[T]{
		Status:  true,
		Message: "Success",
		Result:  v,
	})

	return nil
}

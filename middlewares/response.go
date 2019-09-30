package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func ResponseMiddleware(c *gin.Context) {
	c.Next()
	value, isExist := c.Get("result")
	if !isExist {
		c.JSON(http.StatusBadRequest, Response{
			Code:    http.StatusBadRequest,
			Message: "empty response, please check your parameter",
			Data:    value,
		})
	} else {

	}
}

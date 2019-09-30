package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var router *gin.Engine

func init() {
	router = gin.Default()
	{
		// basic group
		basic := Basic{}
		basicGroup := router.Group("/basic")
		basicGroup.GET("/info", basic.GetBasicInfo)
	}
	{
		// container group
		container := Container{}
		basicGroup := router.Group("/container")
		basicGroup.GET("/list", container.GetContainerList)
	}
}

type Response struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func GetRouter() *gin.Engine {
	return router
}

const (
	STATUS_OK = iota
	STATUS_FAILED
)

var statusMap =map[int]string{
	STATUS_OK: "success",
	STATUS_FAILED: "unknow error",
}

func Result(c *gin.Context, code int, data interface{}, err error) {
	if err != nil {
		// 保存错误
	}
	message, isExist := statusMap[code]
	if !isExist {
		message = statusMap[STATUS_FAILED]
	}
	if code == 0 {
		c.AbortWithStatusJSON(http.StatusOK, Response{
			Code:    code,
			Message: message,
			Data:    data,
		})
		return
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

package apiscommon

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/errors"
)

// SendError 对api返回的错误进行封装
func SendError(c *gin.Context, serviceID string, err error) {
	wrappedErr, ok := err.(*errors.Error)
	if !ok {
		c.JSON(http.StatusInternalServerError, errors.InternalServerError(serviceID, err.Error()))
		return
	}
	c.JSON(int(wrappedErr.Code), wrappedErr.Detail)
}

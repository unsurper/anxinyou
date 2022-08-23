package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//封装http返回
/*
{
	code: 20001,
	data: xxx,
	msg: xx,
}
*/
func Response(c *gin.Context, httpStatus int, code int, data gin.H, msg string) {
	c.JSON(httpStatus, gin.H{"code": code, "data": data, "msg": msg})
}
func Success(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 200, data, msg)
}
func Fail(c *gin.Context, data gin.H, msg string) {
	Response(c, http.StatusOK, 400, data, msg)
}

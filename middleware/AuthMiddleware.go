package middleware

import (
	"anxinyou/common"
	"anxinyou/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取AuthMiddleware header
		tokenString := c.GetHeader("Authorization")
		//验证一下格式
		fmt.Println(tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限1不足"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParesToken(tokenString)
		if err != nil || !token.Valid { //如果token是无效的
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限2不足"})
			log.Println(err)
			c.Abort()
			return
		}

		//验证通过后获取claim 中的userID
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		//用户
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "查询不到用户ID"})
			c.Abort()
			return
		}
		//用户存在 将user 的信息写入上下文
		c.Set("user", user)
		c.Next()
	}
}

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取AuthMiddleware header
		tokenString := c.GetHeader("Authorization")
		//验证一下格式
		fmt.Println(tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限1不足"})
			c.Abort()
			return
		}
		tokenString = tokenString[7:]

		token, claims, err := common.ParesToken(tokenString)
		if err != nil || !token.Valid { //如果token是无效的
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限2不足"})
			log.Println(err)
			c.Abort()
			return
		}

		//验证通过后获取claim 中的userID
		userId := claims.UserId
		DB := common.GetDB()
		var user model.Admin
		DB.First(&user, userId)

		//用户
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "查询不到用户ID"})
			c.Abort()
			return
		}
		//用户存在 将user 的信息写入上下文
		c.Set("admin", user)
		c.Next()
	}
}

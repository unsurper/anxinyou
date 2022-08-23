package middleware

import (
	"anxinyou/common"
	"anxinyou/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
)

func FilterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//获取AuthMiddleware header
		title := c.PostForm("title")
		//验证一下格式
		DB := common.GetDB()

		var filter model.Filter
		DB.Table("filters").Last(&filter)
		fmt.Println(filter)

		var filters []model.Filter
		DB.Table("filters").Find(&filters)
		fmt.Println(filters)

		for i := 0; i < len(filters); i++ {
			judge, _ := regexp.MatchString(filters[i].Word, title)
			fmt.Println(judge, filters[i].Word)
			if judge == true {
				c.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "标题含违规词禁止发布"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

package controller

import (
	"anxinyou/common"
	"anxinyou/model"
	"anxinyou/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserUpPost(c *gin.Context) {
	DB := common.GetDB()
	upuser := c.PostForm("upuser")
	uppost := c.PostForm("uppost")

	var up model.Up
	DB.Table("ups").Where("upuser = ? and uppost = ?", upuser, uppost).First(&up)
	if up.ID != 0 {
		response.Fail(c, nil, "不能重复点赞")
		return
	}

	var post model.Post
	DB.Table("posts").Where("id = ?", uppost).Find(&post)
	DB.Model(&post).Update("uppost", gorm.Expr("uppost + ?", 1)) //点赞加一

	newUp := model.Up{
		Upuser: upuser,
		Uppost: uppost,
	}
	DB.Create(&newUp)
	response.Success(c, nil, "点赞成功")
}

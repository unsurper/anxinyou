package controller

import (
	"anxinyou/common"
	"anxinyou/model"
	"anxinyou/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateComment(c *gin.Context) {
	DB := common.GetDB()
	telephone := c.PostForm("telephone")
	text := c.PostForm("text")
	postid := c.PostForm("postid")

	var user model.User
	DB.Table("users").Where("telephone = ?", telephone).First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	if text == " " {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "评论不能为空")
		return
	}

	newComment := model.Comment{
		Telephone: telephone,
		Text:      text,
		Postid:    postid,
	}
	DB.Create(&newComment)
	//返回结果
	response.Success(c, nil, "评论成功")
}

func GetComment(c *gin.Context) {
	DB := common.GetDB()
	id := c.Query("ID")
	var comment []model.Comment
	fmt.Println("ID:" + id)
	DB.Table("comments").Where("postid = ?", id).Find(&comment)
	response.Success(c, gin.H{"data": comment}, "取出评论")
}

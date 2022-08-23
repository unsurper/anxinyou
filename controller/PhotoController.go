package controller

import (
	"anxinyou/common"
	"anxinyou/model"
	"anxinyou/response"
	"anxinyou/util"
	"github.com/gin-gonic/gin"
)

func UploadPhoto(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传文件出错")
	}
	//c.SaveUploadedFile(file,file.Filename)
	file.Filename = util.RandomString(10)
	c.SaveUploadedFile(file, "./static/photo/"+file.Filename+".jpg")
	c.String(200, "?imageName=F:/goCode/anxinyou/static/photo/"+file.Filename+".jpg")
}

func UploadCover(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传文件出错")
	}
	//c.SaveUploadedFile(file,file.Filename)
	file.Filename = util.RandomString(10)
	c.SaveUploadedFile(file, "./static/photo/"+file.Filename+".jpg")
	c.String(200, "/photo/"+file.Filename+".jpg")
}

func UploadPortrait(c *gin.Context) {
	DB := common.GetDB()
	telephone := c.PostForm("telephone")
	portrait := c.PostForm("portrait")
	//c.SaveUploadedFile(file,file.Filename)
	DB.Table("users").Where("telephone = ?", telephone).Update("portrait", portrait)
	response.Success(c, nil, "上传头像成功")
}

func GetPhoto(c *gin.Context) {
	imageName := c.Query("imageName")
	c.File(imageName)
}
func GetmyPortrait(c *gin.Context) {
	DB := common.GetDB()
	telephone := c.Query("telephone")
	var user model.User
	DB.Table("users").Where("telephone = ?", telephone).First(&user)
	response.Success(c, gin.H{"data": user.Portrait}, "取得头像成功")
}

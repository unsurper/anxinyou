package controller

import (
	"anxinyou/common"
	"anxinyou/model"
	"anxinyou/response"
	"github.com/gin-gonic/gin"
	"math/rand"
	"time"
)

func Createfilter(c *gin.Context) {
	DB := common.GetDB()
	word := c.PostForm("word")

	newfilter := model.Filter{
		Word: word,
	}
	DB.Create(&newfilter)
	response.Success(c, nil, "创建违规词成功")
}

func Suijisu(c *gin.Context) {
	DB := common.GetDB()

	var post []model.Post
	DB.Find(&post)

	for i := 0; i < len(post); i++ {
		rand.Seed(time.Now().UnixNano())
		b := rand.Intn(10000)
		u := rand.Intn(1000)
		DB.Model(&post[i]).Update("browsepost", b) //浏览加一
		DB.Model(&post[i]).Update("uppost", u)     //浏览加一
	}
	var comment []model.Comment
	DB.Find(&comment)
	for i := 0; i < len(comment); i++ {
		rand.Seed(time.Now().UnixNano())
		u := rand.Intn(100)
		DB.Model(&comment[i]).Update("upcomment", u) //浏览加一
	}
	response.Success(c, nil, "动态浏览与点赞成功")
}

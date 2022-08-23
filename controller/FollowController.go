package controller

import (
	"anxinyou/common"
	"anxinyou/model"
	"anxinyou/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FollowUser(c *gin.Context) {
	DB := common.GetDB()
	username := c.PostForm("username")
	follower := c.PostForm("follower")

	if username == follower {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "不能关注自己")
		return
	}

	var follow model.Follow
	DB.Table("follows").Where("username = ? and follower = ?", username, follower).First(&follow)
	if follow.ID != 0 {
		response.Fail(c, nil, "不能重复关注")
		return
	}
	newFollow := model.Follow{
		Username: username,
		Follower: follower,
	}
	DB.Create(&newFollow)
	response.Success(c, nil, "关注成功")
}

func GetFollowuser(c *gin.Context) {
	DB := common.GetDB()
	follower := c.Query("telephone")
	var follow []model.Follow
	DB.Table("follows").Where("follower = ?", follower).Find(&follow)

	var user []model.User
	var users []model.User
	for i := 0; i < len(follow); i++ {
		DB.Table("users").Where("telephone = ?", follow[i].Username).Find(&user)
		users = append(users, user...)
	}

	response.Success(c, gin.H{"data": users}, "取得我的关注列表")

}

func IFFollow(c *gin.Context) {
	DB := common.GetDB()
	follower := c.Query("follower")
	username := c.Query("username")
	var follow model.Follow
	DB.Table("follows").Where("follower = ? and username = ?", follower, username).First(&follow)
	if follow.ID != 0 {
		response.Success(c, gin.H{"data": true}, "已关注")
		return
	} else {
		response.Success(c, gin.H{"data": false}, "未关注")
	}
}

func NoFollow(c *gin.Context) {
	DB := common.GetDB()
	follower := c.PostForm("follower")
	username := c.PostForm("username")
	var follow model.Follow
	DB.Unscoped().Table("follows").Where("follower = ? and username = ?", follower, username).Delete(&follow)
	response.Success(c, nil, "取消关注")
}

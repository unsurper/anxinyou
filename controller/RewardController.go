package controller

import (
	"anxinyou/common"
	"anxinyou/model"
	"anxinyou/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func PostReward(c *gin.Context) {
	DB := common.GetDB()

	sender := c.PostForm("sender")
	receiver := c.PostForm("receiver")
	value := c.PostForm("value")

	var user model.User
	DB.Table("users").Where("telephone = ?", sender).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "发送者用户不存在")
		return
	}
	money, _ := strconv.ParseFloat(value, 64)
	if money == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 423, nil, "打赏金额不能为0")
		return
	}
	if user.Money < money {
		response.Response(c, http.StatusUnprocessableEntity, 424, nil, "你的余额不足")
		return
	}

	DB.Model(&user).Update("money", gorm.Expr("money - ?", money)) //扣钱

	var userreceiver model.User
	DB.Table("users").Where("telephone = ?", receiver).First(&userreceiver)
	DB.Model(&userreceiver).Update("money", gorm.Expr("money + ?", money)) //加钱

	response.Success(c, gin.H{"msg": "打赏成功"}, "打赏成功")
}

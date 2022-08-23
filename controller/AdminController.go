package controller

import (
	"anxinyou/common"
	"anxinyou/model"
	"anxinyou/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func AdminRegister(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	//var requestUser = model.User{}
	//c.Bind(&requestUser)

	name := c.PostForm("name")
	password := c.PostForm("password")
	//数据验证
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//如果名称没有传，给一个10位的字符随机串
	if len(name) == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 423, nil, "用户名不能为空")
		return
	}
	//判断手机号是否存在
	if isNameExist(DB, name) {
		response.Response(c, http.StatusUnprocessableEntity, 424, nil, "用户已经存在")
		return
	}
	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}
	newAdmin := model.Admin{
		Name:     name,
		Password: string(hasedPassword),
		Auth:     0,
	}
	DB.Create(&newAdmin)
	//返回结果
	response.Success(c, nil, "注册成功")
}

func AdminLogin(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	name := c.PostForm("name")
	password := c.PostForm("password")
	//数据验证
	if len(name) == 0 {
		response.Response(c, http.StatusOK, 423, nil, "用户名不能为空")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusOK, 422, nil, "密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var admin model.Admin
	DB.Table("admins").Where("name = ?", name).First(&admin)
	fmt.Println(name, password, admin)
	if admin.ID == 0 {
		response.Response(c, http.StatusOK, 200, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusOK, 400, nil, "密码错误")
		return
	}
	//发放token给前端
	token, err := common.AdminReleaseToken(admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate err : %v", err)
		return
	}
	//返回结果

	response.Success(c, gin.H{"token": token}, "登入成功")
}
func AdminInfo(c *gin.Context) {
	user, _ := c.Get("admin")
	//dto.ToUserDto()只返回dto里面的数据
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": user}})
}

func isNameExist(db *gorm.DB, name string) bool {
	var user model.Admin
	db.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

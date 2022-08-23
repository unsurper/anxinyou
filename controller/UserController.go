package controller

import (
	"anxinyou/common"
	"anxinyou/dto"
	"anxinyou/model"
	"anxinyou/response"
	"anxinyou/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	//var requestUser = model.User{}
	//c.Bind(&requestUser)

	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {

		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//如果名称没有传，给一个10位的字符随机串
	if len(name) == 0 || name == "undefined" {
		name = util.RandomString(10)
	}
	//判断手机号是否存在
	if isTelephoneExist(DB, telephone) {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")
		return
	}
	//创建用户
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		return
	}
	newUser := model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hasedPassword),
	}
	DB.Create(&newUser)
	//返回结果
	response.Success(c, nil, "注册成功")
}

func Login(c *gin.Context) {
	DB := common.GetDB()
	//获取参数
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		return
	}
	//判断手机号是否存在
	var user model.User
	DB.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}
	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 400, nil, "密码错误")
		return
	}
	//发放token给前端
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate err : %v", err)
		return
	}
	//返回结果

	response.Success(c, gin.H{"token": token}, "登入成功")
}

func GetPostUserbytelephone(c *gin.Context) {
	DB := common.GetDB()
	telephone := c.Query("telephone")
	var user model.User
	DB.Table("users").Where("telephone = ?", telephone).Find(&user)
	//fmt.Println(user)
	response.Success(c, gin.H{"data": user}, "查找用户成功")
}

func GetAllUser(c *gin.Context) {
	DB := common.GetDB()
	//当前页
	current, _ := strconv.Atoi(c.Query("current"))
	//每页大小
	size, _ := strconv.Atoi(c.Query("size"))
	if current == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "当前页不能为空")
		return
	}
	if size == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "页面个数不能为空")
		return
	}

	var user []model.User
	fmt.Println("current", current, "size", size)
	offset := (current - 1) * size
	if err := DB.Order("id DESC").Offset(offset).Limit(size).Find(&user).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    502,
			"message": "查询数据异常",
		})
		return
	}
	var total int64
	DB.Model(&model.User{}).Count(&total)

	//fmt.Println(user)
	response.Success(c, gin.H{"data": user, "current": current, "size": size, "total": total}, "查找用户成功")
}

func DeletedUser(c *gin.Context) {
	DB := common.GetDB()
	//当前页
	userid, _ := strconv.Atoi(c.PostForm("id"))
	var user model.User
	DB.Unscoped().Table("users").Where("ID = ?", userid).Delete(&user)
	//fmt.Println(user)
	response.Success(c, nil, "用户删除成功")
}

func Info(c *gin.Context) {
	user, _ := c.Get("user")
	//dto.ToUserDto()只返回dto里面的数据
	c.JSON(http.StatusOK, gin.H{"code": 200, "data": gin.H{"user": dto.ToUserDto(user.(model.User))}})
}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

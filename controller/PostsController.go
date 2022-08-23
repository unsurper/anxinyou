package controller

import (
	"anxinyou/common"
	"anxinyou/model"
	"anxinyou/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func CreatePost(c *gin.Context) {
	DB := common.GetDB()
	telephone := c.PostForm("telephone")
	title := c.PostForm("title")
	text := c.PostForm("text")
	cover := c.PostForm("cover")

	var user model.User
	DB.Table("users").Where("telephone = ?", telephone).First(&user)
	// SELECT * FROM users WHERE name = 'jinzhu' ORDER BY id LIMIT 1;
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")
		return
	}

	newPost := model.Post{
		Telephone: telephone,
		Text:      text,
		Title:     title,
		Cover:     cover,
	}
	DB.Create(&newPost)
	//返回结果
	response.Success(c, nil, "发布帖子成功")
}
func GetPost(c *gin.Context) {
	DB := common.GetDB()
	var post []model.Post
	//DB.Scopes(util.Paginate(1,100)).Find(&post)
	DB.Find(&post)
	response.Success(c, gin.H{"data": post}, "提取帖子成功")
}

func Getpostbyid(c *gin.Context) {
	DB := common.GetDB()
	id := c.Query("ID")
	var post model.Post
	DB.Table("posts").Where("id = ?", id).Find(&post)
	DB.Model(&post).Update("browsepost", gorm.Expr("browsepost + ?", 1)) //浏览加一
	response.Success(c, gin.H{"data": post}, "打开帖子成功")
}

func Getpostbyfollow(c *gin.Context) {
	DB := common.GetDB()
	follower := c.Query("follower")
	var follow []model.Follow
	DB.Table("follows").Where("follower = ?", follower).Find(&follow)

	var post []model.Post
	var postsum []model.Post
	for i := 0; i < len(follow); i++ {
		DB.Table("posts").Where("telephone = ?", follow[i].Username).Find(&post)
		postsum = append(postsum, post...)
	}
	response.Success(c, gin.H{"data": postsum}, "取得关注帖子成功")
}

func Getpostbymypost(c *gin.Context) {
	DB := common.GetDB()
	telephone := c.Query("telephone")
	var post []model.Post
	DB.Table("posts").Where("telephone = ?", telephone).Find(&post)
	response.Success(c, gin.H{"data": post}, "取得关注帖子成功")
}
func GetAllPost(c *gin.Context) {
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

	var post []model.Post
	fmt.Println("current", current, "size", size)
	offset := (current - 1) * size

	if err := DB.Order("id DESC").Offset(offset).Limit(size).Find(&post).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    502,
			"message": "查询数据异常",
		})
		return
	}

	var total int64
	DB.Model(&model.Post{}).Count(&total)

	//fmt.Println(user)
	response.Success(c, gin.H{"data": post, "current": current, "size": size, "total": total}, "查找动态成功")
}

func DeletedPost(c *gin.Context) {
	DB := common.GetDB()
	//当前页
	postid, _ := strconv.Atoi(c.PostForm("id"))
	var post model.Post
	DB.Unscoped().Table("posts").Where("ID = ?", postid).Delete(&post)
	//fmt.Println(user)
	response.Success(c, nil, "帖子删除成功")
}

func UserDeletedPost(c *gin.Context) {
	DB := common.GetDB()
	//当前页
	postid := c.PostForm("postid")
	var post model.Post
	if postid == "" {
		response.Fail(c, nil, "查询不到帖子信息")
		return
	}

	DB.Unscoped().Table("posts").Where("ID = ?", postid).Delete(&post)
	//fmt.Println(user)
	response.Success(c, nil, "帖子删除成功")
}

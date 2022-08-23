package main

import (
	"anxinyou/controller"
	"anxinyou/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	//r.Use(middleware.AdminCORSMiddleware())

	r.GET("/anxinyou/sys/info", middleware.AdminAuthMiddleware(), controller.AdminInfo)                //带token的中间件取Info信息
	r.POST("/anxinyou/sys/register", controller.AdminRegister)                                         //注册
	r.POST("/anxinyou/sys/login", controller.AdminLogin)                                               //登入
	r.GET("/anxinyou/sys/getalluser", middleware.AdminAuthMiddleware(), controller.GetAllUser)         //带token的中间件取所有用户信息
	r.GET("/anxinyou/sys/getallcontract", middleware.AdminAuthMiddleware(), controller.GetAllContract) //带token的中间件取所有何荣信息
	r.GET("/anxinyou/sys/getallpost", middleware.AdminAuthMiddleware(), controller.GetAllPost)         //带token的中间件取所有动态信息
	r.POST("/anxinyou/sys/checkcontract", middleware.AdminAuthMiddleware(), controller.CheckContract)  //带token的中间件审核合同
	r.POST("/anxinyou/sys/deleteduser", middleware.AdminAuthMiddleware(), controller.DeletedUser)      //带token的中间件删除用户
	r.POST("/anxinyou/sys/deletedpost", middleware.AdminAuthMiddleware(), controller.DeletedPost)      //带token的中间件删除动态

	r.POST("/api/auth/register", controller.Register)                     //注册
	r.POST("/api/auth/login", controller.Login)                           //登入
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info) //带token的中间件取Info信息

	r.POST("/api/auth/createpost", middleware.FilterMiddleware(), controller.CreatePost) //创建帖子post方法
	r.POST("/api/auth/uploadphoto", controller.UploadPhoto)                              //上传图片post方法
	r.POST("/api/auth/uploadcover", controller.UploadCover)                              //上传帖子封面图片post方法

	r.GET("/api/auth/getphoto", controller.GetPhoto)                             //取图片
	r.GET("/api/auth/getpost", controller.GetPost)                               //取帖子
	r.GET("/api/auth/getpostbyid", controller.Getpostbyid)                       //取帖子通过id
	r.GET("/api/auth/getpostbyfollow", controller.Getpostbyfollow)               //取帖子通过关注列表
	r.GET("/api/auth/getpostbymypost", controller.Getpostbymypost)               //取自己发布的帖子
	r.GET("/api/auth/getpostuserbytelephone", controller.GetPostUserbytelephone) //取user通过电话
	r.GET("/api/auth/getcommentbyid", controller.GetComment)                     //取评论
	r.POST("/api/auth/createcomment", controller.CreateComment)                  //发布评论方法
	r.POST("/api/auth/followuser", controller.FollowUser)                        //关注
	r.POST("/api/auth/useruppost", controller.UserUpPost)                        //点赞
	r.GET("/api/auth/getfollowuser", controller.GetFollowuser)                   //取得关注列表
	r.POST("/api/auth/uploadportrait", controller.UploadPortrait)                //上传头像
	r.GET("/api/auth/getmyportrait", controller.GetmyPortrait)                   //得到头像
	r.POST("/api/auth/sendchat", controller.SendChat)                            //发送消息
	r.GET("/api/auth/listchat", controller.ListChat)                             //得到聊天信息
	r.GET("/api/auth/getroomchat", controller.Getroomchat)                       //得到聊天信息
	r.GET("/api/auth/iffollow", controller.IFFollow)                             //判断帖子是否被关注
	r.POST("/api/auth/nofollow", controller.NoFollow)                            //判断帖子是否被关注
	r.POST("/api/auth/createfilter", controller.Createfilter)                    //判断帖子是否被关注
	r.POST("/api/auth/postreward", controller.PostReward)                        //打赏功能
	r.POST("/api/auth/deletedpost", controller.UserDeletedPost)                  //删除帖子

	r.POST("/api/auth/createcontract", controller.CreateContract)         //创建合同
	r.GET("/api/auth/getcontractfirst", controller.GetContractFirst)      //获取第一状态待接受合同
	r.GET("/api/auth/getcontractsecond", controller.GetContractSecond)    //获取第二状态待编辑合同
	r.POST("/api/auth/paycontract", controller.PayContract)               //支付合同
	r.POST("/api/auth/editcontractsecond", controller.EditContractSecond) //编辑合同详情
	r.GET("/api/auth/getcontractthird", controller.GetContractThird)      //获取第三状态待三方确定合同
	r.POST("/api/auth/ifokcontract", controller.IfOkContract)             //支付合同
	r.GET("/api/auth/getcontractforth", controller.GetContractForth)      //获取第四状态进行中合同
	r.GET("/api/auth/getcontractfifth", controller.GetContractFifth)      //获取第五状态完成后合同

	r.POST("/api/auth/suijishu", controller.Suijisu) //随机数

	return r
}

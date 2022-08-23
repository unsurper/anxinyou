package controller

import (
	"anxinyou/common"
	"anxinyou/model"
	"anxinyou/response"
	"anxinyou/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendChat(c *gin.Context) {
	DB := common.GetDB()
	sender := c.PostForm("sender")
	receiver := c.PostForm("receiver")
	content := c.PostForm("content")

	var user model.User
	DB.Table("users").Where("telephone = ?", sender).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "发送者用户不存在")
		return
	}
	DB.Table("users").Where("telephone = ?", receiver).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "接受者用户不存在")
		return
	}

	var chatroom model.Chatroom

	DB.Table("chatrooms").Where("usera = ? and userb = ?", sender, receiver).First(&chatroom)
	if chatroom.ID != 0 {
		DB.Table("chatrooms").Where("ID = ?", chatroom.ID).Update("newmsg", content)
		DB.Table("chatrooms").Where("ID = ?", chatroom.ID).Update("who", sender)
	} else {
		DB.Table("chatrooms").Where("userb = ?", sender).First(&chatroom)
		if chatroom.ID != 0 {
			DB.Table("chatrooms").Where("ID = ?", chatroom.ID).Update("newmsg", content)
			DB.Table("chatrooms").Where("ID = ?", chatroom.ID).Update("who", sender)
		} else {
			newchatroom := model.Chatroom{
				Newmsg: content,
				Who:    sender,
				Usera:  sender,
				Userb:  receiver,
			}
			DB.Create(&newchatroom)
		}
	}

	newChat := model.Chat{
		Sender:   sender,
		Receiver: receiver,
		Content:  content,
	}
	DB.Create(&newChat)
	response.Success(c, nil, "发送消息成功")
}

func ListChat(c *gin.Context) {
	DB := common.GetDB()

	sender := c.Query("sender")
	receiver := c.Query("receiver")

	var chat1 []model.Chat
	var chat2 []model.Chat
	var chatsum []model.Chat

	DB.Table("chats").Where("sender = ? and receiver = ?", sender, receiver).Find(&chat1)
	DB.Table("chats").Where("sender = ? and receiver = ?", receiver, sender).Find(&chat2)

	if len(chat1) == 0 && len(chat2) == 0 {
		response.Success(c, nil, "双方暂无沟通消息")
		return
	}

	chatsum = append(chatsum, chat1...)
	chatsum = append(chatsum, chat2...)

	util.SortByID(chatsum)
	response.Success(c, gin.H{"data": chatsum}, "成功返回沟通消息")
}

func Getroomchat(c *gin.Context) {
	DB := common.GetDB()
	user := c.Query("telephone")
	var chatroom []model.Chatroom
	var chatroomsum []model.Chatroom
	var userinfo []model.User
	var users []model.User

	DB.Table("chatrooms").Where("usera = ?", user).Find(&chatroom)
	chatroomsum = append(chatroomsum, chatroom...)
	DB.Table("chatrooms").Where("userb = ?", user).Find(&chatroom)
	chatroomsum = append(chatroomsum, chatroom...)
	for i := 0; i < len(chatroomsum); i++ {
		if user == chatroomsum[i].Usera {
			DB.Table("users").Where("telephone = ?", chatroomsum[i].Userb).First(&userinfo)
			users = append(users, userinfo...)
		} else {
			DB.Table("users").Where("telephone = ?", chatroomsum[i].Usera).First(&userinfo)
			users = append(users, userinfo...)
		}
	}

	response.Success(c, gin.H{"data": chatroomsum, "other": users}, "成功返回最新消息")
}

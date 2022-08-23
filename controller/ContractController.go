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
	"time"
)

func CreateContract(c *gin.Context) {
	DB := common.GetDB()

	business := c.PostForm("business")   //合同发起者
	buyers := c.PostForm("buyers")       //合同接收者
	offer := c.PostForm("offer")         //合同发起者报价
	guarantee := c.PostForm("guarantee") //发起者担保价格

	if business == buyers {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "不能给自己发起合同")
		return
	}

	var user model.User
	DB.Table("users").Where("telephone = ?", business).First(&user)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 423, nil, "合同发起者用户不存在")
		return
	}
	money, _ := strconv.ParseFloat(guarantee, 64)
	if user.Money < money {
		response.Response(c, http.StatusUnprocessableEntity, 424, nil, "你的余额不足担保")
		return
	}
	DB.Model(&user).Update("money", gorm.Expr("money - ?", money)) //扣钱

	var userreceiver model.User
	DB.Table("users").Where("telephone = ?", buyers).First(&userreceiver)
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 425, nil, "用户不存在请检查手机号是否正确")
		return
	}
	floatoffer, _ := strconv.ParseFloat(offer, 64)
	floatguarantee, _ := strconv.ParseFloat(guarantee, 64)
	newContract := model.Contract{
		Business:  business,
		Buyers:    buyers,
		Offer:     floatoffer,
		Guarantee: floatguarantee,
	}
	DB.Create(&newContract)

	response.Success(c, nil, "发起合同成功")
}

func GetContractFirst(c *gin.Context) {
	DB := common.GetDB()

	telephone := c.Query("telephone")
	var contractlist []model.Contract
	DB.Table("contracts").Where("buyers = ?", telephone).Find(&contractlist)

	var user []model.User
	var users []model.User
	for i := 0; i < len(contractlist); i++ {
		DB.Table("users").Where("telephone = ?", contractlist[i].Business).Find(&user)
		users = append(users, user...)
	}
	response.Success(c, gin.H{"contract": contractlist, "business": users}, "带接收的合同列表")
}

func GetContractSecond(c *gin.Context) {
	DB := common.GetDB()

	telephone := c.Query("telephone")
	var contractlist []model.Contractfirst
	DB.Table("contractfirsts").Where("business = ? and state <= ?", telephone, 1).Find(&contractlist)

	var user []model.User
	var users []model.User
	for i := 0; i < len(contractlist); i++ {
		DB.Table("users").Where("telephone = ?", contractlist[i].Buyers).Find(&user)
		users = append(users, user...)
	}
	response.Success(c, gin.H{"contract": contractlist, "buyers": users}, "取待处理的合同列表")
}

func EditContractSecond(c *gin.Context) {
	DB := common.GetDB()

	contractid := c.PostForm("contractid")
	starttime := c.PostForm("starttime")
	endtime := c.PostForm("endtime")
	contractimg := c.PostForm("contractimg")

	var contractlist model.Contractfirst
	DB.Table("contractfirsts").Where("ID = ? ", contractid).First(&contractlist)
	DB.Table("contractfirsts").Where("ID = ? ", contractid).Update("starttime", starttime)
	DB.Table("contractfirsts").Where("ID = ? ", contractid).Update("endtime", endtime)
	DB.Table("contractfirsts").Where("ID = ? ", contractid).Update("contractimg", contractimg)
	if contractlist.State == 0 {
		DB.Model(&contractlist).Update("state", gorm.Expr("state + ?", 1))
	}

	response.Success(c, nil, "编辑上传成功")
}

func GetContractThird(c *gin.Context) {
	DB := common.GetDB()

	telephone := c.Query("telephone")
	var contractlist []model.Contractfirst
	DB.Table("contractfirsts").Where("buyers = ? and state >= ? and state <= ?", telephone, 1, 2).Find(&contractlist)

	var user []model.User
	var users []model.User
	for i := 0; i < len(contractlist); i++ {
		DB.Table("users").Where("telephone = ?", contractlist[i].Business).Find(&user)
		users = append(users, user...)
	}
	response.Success(c, gin.H{"contract": contractlist, "business": users}, "取待确定的合同列表")
}

func GetContractForth(c *gin.Context) {
	DB := common.GetDB()

	telephone := c.Query("telephone")
	var contractlist []model.Contractfirst
	var contractsum []model.Contractfirst
	DB.Table("contractfirsts").Where("buyers = ? and state = ?", telephone, 3).Find(&contractlist)

	nowtime := time.Now().Unix()
	for i := 0; i < len(contractlist); i++ {

		tim := contractlist[i].Endtime.Unix()

		if tim > nowtime {
			contractsum = append(contractsum, contractlist[i])
		}
	}
	var user []model.User
	var users []model.User
	for i := 0; i < len(contractsum); i++ {
		DB.Table("users").Where("telephone = ?", contractsum[i].Business).Find(&user)
		users = append(users, user...)
	}
	response.Success(c, gin.H{"contract": contractsum, "business": users}, "取进行中的合同列表")
}

func GetContractFifth(c *gin.Context) {
	DB := common.GetDB()

	telephone := c.Query("telephone")
	var contractlist []model.Contractfirst
	var contractsum []model.Contractfirst
	DB.Table("contractfirsts").Where("buyers = ? and state = ?", telephone, 3).Find(&contractlist)

	nowtime := time.Now().Unix()
	for i := 0; i < len(contractlist); i++ {

		tim := contractlist[i].Endtime.Unix()
		if tim < nowtime {
			contractsum = append(contractsum, contractlist[i])
		}
	}
	var user []model.User
	var users []model.User
	for i := 0; i < len(contractsum); i++ {
		DB.Table("users").Where("telephone = ?", contractsum[i].Business).Find(&user)
		users = append(users, user...)
	}
	response.Success(c, gin.H{"contract": contractsum, "business": users}, "取进行中的合同列表")
}

func PayContract(c *gin.Context) {
	DB := common.GetDB()

	telephone := c.PostForm("buyers")
	contractid := c.PostForm("contractid")
	ifpay := c.PostForm("ifpay")
	if ifpay == "0" {
		var contract model.Contract
		DB.Unscoped().Table("contracts").Where("ID = ?", contractid).Delete(&contract)
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "已拒绝")
		return
	}

	var contract model.Contract
	DB.Table("contracts").Where("ID = ?", contractid).Find(&contract)

	var user model.User
	DB.Table("users").Where("telephone = ?", telephone).First(&user)

	fmt.Println(user, user.Money, "要支付", contract.Offer)
	if user.Money < contract.Offer {
		response.Response(c, http.StatusUnprocessableEntity, 423, nil, "余额不足请充值")
		return
	}
	DB.Model(&user).Update("money", gorm.Expr("money - ?", contract.Offer)) //扣钱

	newContractfirst := model.Contractfirst{
		Business:  contract.Business,
		Buyers:    contract.Buyers,
		Offer:     contract.Offer,
		Guarantee: contract.Guarantee,
		Starttime: time.Now(),
		Endtime:   time.Now(),
		State:     0,
	}
	DB.Create(&newContractfirst)
	DB.Unscoped().Table("contracts").Where("ID = ?", contractid).Delete(&contract)

	response.Success(c, nil, "支付成功")
}

func IfOkContract(c *gin.Context) {
	DB := common.GetDB()

	telephone := c.PostForm("buyers")
	contractid := c.PostForm("contractid")
	ifok := c.PostForm("ifok")
	var user model.User
	DB.Table("users").Where("telephone = ?", telephone).First(&user)
	var contractfirst model.Contractfirst
	DB.Table("contractfirsts").Where("ID = ?", contractid).Find(&contractfirst)
	if ifok == "0" {
		money := contractfirst.Offer
		DB.Unscoped().Table("contractfirsts").Where("ID = ?", contractid).Delete(&contractfirst)
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "取消订单")
		DB.Model(&user).Update("money", gorm.Expr("money + ?", money)) //加钱
		return
	}
	if user.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 423, nil, "用户不存在")
		return
	}
	if contractfirst.State == 1 {
		DB.Model(&contractfirst).Update("state", gorm.Expr("state + ?", 1))
		response.Success(c, nil, "已确认请等待三方平台审批")
		return
	}
	response.Success(c, nil, "请勿重复确认")
}

func CheckContract(c *gin.Context) {
	DB := common.GetDB()

	contractid := c.PostForm("contractid")
	ifok := c.PostForm("ifok")
	fmt.Println(contractid, ifok)
	var contractfirst model.Contractfirst
	DB.Table("contractfirsts").Where("ID = ?", contractid).Find(&contractfirst)
	if contractfirst.ID == 0 {
		response.Response(c, http.StatusUnprocessableEntity, 423, nil, "合同不存在")
		return
	}
	if contractfirst.State == 2 && ifok == "1" {
		DB.Model(&contractfirst).Update("state", gorm.Expr("state + ?", 1))
		response.Success(c, nil, "审批通过")
		return
	} else if ifok == "0" {
		DB.Model(&contractfirst).Update("state", gorm.Expr("state - ?", 1))
		response.Success(c, nil, "审批未通过")
		return
	} else if contractfirst.State < 2 {
		response.Response(c, http.StatusOK, 423, nil, "该合同不需审批")
		return
	}
	response.Success(c, nil, "请勿重复确认")
}

func GetAllContract(c *gin.Context) {
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
	state, _ := strconv.Atoi(c.Query("state"))

	var contract []model.Contractfirst
	fmt.Println("current", current, "size", size)
	offset := (current - 1) * size
	if state == -1 {
		if err := DB.Order("id DESC").Offset(offset).Limit(size).Find(&contract).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    502,
				"message": "查询数据异常",
			})
			return
		}
	} else {
		if err := DB.Order("id DESC").Where("state = ?", state).Offset(offset).Limit(size).Find(&contract).Error; err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code":    502,
				"message": "查询数据异常",
			})
			return
		}
	}

	var total int64
	DB.Model(&model.Contractfirst{}).Count(&total)

	//fmt.Println(user)
	response.Success(c, gin.H{"data": contract, "current": current, "size": size, "total": total}, "查找合同成功")
}

package common

import (
	"anxinyou/model"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: viper.GetString("datasource.username") +
			":" + viper.GetString("datasource.password") +
			"@tcp(" + viper.GetString("datasource.host") +
			":" + viper.GetString("datasource.port") +
			")/" + viper.GetString("datasource.database") +
			"?charset=" + viper.GetString("datasource.charset") +
			"&parseTime=True&loc=Local",
		DefaultStringSize: 171,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		//GORM 会在事务里执行写入操作（创建、更新、删除）。如果没有这方面的要求，您可以在初始化时禁用它。
		//NamingStrategy: schema.NamingStrategy{
		//	TablePrefix: "gva_",   // 表名前缀，`User` 的表名应该是 `t_users`
		//	SingularTable: false, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		//},
		//GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
		DisableForeignKeyConstraintWhenMigrating: true, //  主张逻辑外键(代码里面自动外键外键关系)
	})
	if err != nil {
		fmt.Println(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           //连接池中最大的空闲连接数
	sqlDB.SetMaxOpenConns(100)          //连接池最多容纳的连接数量
	sqlDB.SetConnMaxLifetime(time.Hour) //连接池中连接的最大可复用时间

	db.AutoMigrate(&model.Admin{})         //建用户表
	db.AutoMigrate(&model.User{})          //建用户表
	db.AutoMigrate(&model.Post{})          //建立帖子表
	db.AutoMigrate(&model.Comment{})       //建立评论表
	db.AutoMigrate(&model.Follow{})        //建立关注表
	db.AutoMigrate(&model.Up{})            //建立关注表
	db.AutoMigrate(&model.Chat{})          //建立聊天消息表
	db.AutoMigrate(&model.Chatroom{})      //建立聊天建立表
	db.AutoMigrate(&model.Filter{})        //建立屏蔽词表
	db.AutoMigrate(&model.Contract{})      //建立合同发起表
	db.AutoMigrate(&model.Contractfirst{}) //建立合同详情表

	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}

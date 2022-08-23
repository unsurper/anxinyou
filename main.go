package main

import (
	"anxinyou/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	InitConfig()
	db := common.InitDB()
	log.Println(db)

	r := gin.Default()
	r = CollectRoute(r)
	r.MaxMultipartMemory = 8 << 20 //8mib 设置最大的上传文件大小

	r.Static("/static", "./static")

	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

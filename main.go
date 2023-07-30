package main

import (
	"fmt"
	logs "github.com/danbai225/go-logs"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"index_Demo/bootstrap/global"
	"index_Demo/config"
	mysql "index_Demo/dao/mysql"
	"index_Demo/dao/redisServer"
	"index_Demo/router"
)

func main() {
	//加载config文件
	err := config.InfoConfig()
	if err != nil {
		logs.Err(err.Error())
		return
	}
	//err = config.LoadConfig()
	//if err != nil {
	//	logs.Err(err)
	//	return
	//}

	//连接sql
	err = mysql.Init()
	if err != nil {
		logs.Err(err.Error())
		return
	}
	//连接redis
	redisServer.Init()
	//初始化Gin
	global.Engine = gin.New()
	//加载路由
	router.Router(global.Engine)
	//http://patorjk.com/software/taag/#p=display&h=0&f=Ogre&t=demo
	logs.Info("\n                 ___               _                      _            ___                         \n  __ _   ___    / __\\  __ _   ___ | | __  ___  _ __    __| |          /   \\  ___  _ __ ___    ___  \n / _` | / _ \\  /__\\// / _` | / __|| |/ / / _ \\| '_ \\  / _` | _____   / /\\ / / _ \\| '_ ` _ \\  / _ \\ \n| (_| || (_) |/ \\/  \\| (_| || (__ |   < |  __/| | | || (_| ||_____| / /_// |  __/| | | | | || (_) |\n \\__, | \\___/ \\_____/ \\__,_| \\___||_|\\_\\ \\___||_| |_| \\__,_|       /___,'   \\___||_| |_| |_| \\___/ \n |___/                                                                                             \n")
	_ = global.Engine.Run(fmt.Sprintf(":%d", viper.GetUint16("server.port")))
}

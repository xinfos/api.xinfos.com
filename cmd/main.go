package main

import (
	"api.xinfos.com/driver"
	"api.xinfos.com/pkg/logger"
	"api.xinfos.com/routers"
	"api.xinfos.com/utils/riot"

	"github.com/gin-gonic/gin"
)

//loadResourcesServer 加载资源服务
func loadResourcesServer() {
	driver.InitDB()    //MySQL 初始化
	driver.InitRedis() //Redis 缓存初始化
	riot.InitEngine()  //Riot 搜索初始化
}

func main() {

	loadResourcesServer()

	logger.InitLogger()

	// gin.SetMode(gin.ReleaseMode)
	// ginpprof.Wrap(g)
	gin.SetMode(gin.DebugMode)

	g := routers.Load(gin.New())

	if err := g.Run(":8082"); err != nil {
		logger.F.Error("service start fail: " + err.Error())
	}
}

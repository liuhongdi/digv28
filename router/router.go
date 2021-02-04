package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/liuhongdi/digv28/controller"
	"github.com/liuhongdi/digv28/global"
	"log"
	"runtime/debug"
)

func Router() *gin.Engine {
	router := gin.Default()
	//处理异常
	router.NoRoute(HandleNotFound)
	router.NoMethod(HandleNotFound)
	router.Use(Recover)
	//cors
	config := cors.DefaultConfig()
	//config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	config.AllowOrigins = []string{"http://127.0.0.1","http://google.com", "http://facebook.com"}
	router.Use(cors.New(config))

	// 路径映射:index
	indexc:=controller.NewIndexController()
	router.GET("/index/index", indexc.Index);

	return router
}

func HandleNotFound(c *gin.Context) {
	global.NewResult(c).Error(404,"资源未找到")
	return
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			//打印错误堆栈信息
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			global.NewResult(c).Error(500,"服务器内部错误")
		}
	}()
	//加载完 defer recover，继续后续接口调用
	c.Next()
}
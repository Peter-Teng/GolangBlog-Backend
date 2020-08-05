package main

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/config"
	"MarvelousBlog-Backend/middleware"
	r "MarvelousBlog-Backend/router"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title PP同学个人博客接口文档
// @version 0.1
// @description MarvelousBlog-Backend Swagger接口文档
// @contact.name PP同学
// @contact.email 710955321@qq.com
// @host localhost:8600
func main() {
	gin.SetMode(config.AppMode)
	engine := gin.New()

	//添加日志中间件
	engine.Use(middleware.LoggingMiddleware())

	//添加错误恢复中间件
	engine.Use(gin.Recovery())

	//加载各类router
	r.LoadVisitorRouters(engine)

	//启动swagger
	config.InitSwaggerRouter(engine)

	server := &http.Server{Addr: config.ServerPort, Handler: engine}

	//启动服务
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Service run failed! errMsg = ", err)
			closeResources()
		}
	}()

	//优雅关闭服务(5秒延迟)
	exit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit
	config.Log.Infof(common.SYSTEM_INFO_LOG, "Shutting Down Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	//关闭资源
	defer closeResources()

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Server Shutdown Error Encountered:", err)
	}

	select {
	case <-ctx.Done():
		config.Log.Infof(common.SYSTEM_INFO_LOG, "timeout of 3 seconds.")
	}

}

func closeResources() {
	config.Log.Infof(common.SYSTEM_INFO_LOG, "Resources Closing")
	if err := config.RedisPool.Close(); err != nil {
		config.Log.Error(common.SYSTEM_ERROR_LOG, "Redis pool closed error! errMsg = ", err)
	}
}

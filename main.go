package main

import (
	"MarvelousBlog-Backend/common"
	"MarvelousBlog-Backend/config"
	"MarvelousBlog-Backend/entity/model"
	"MarvelousBlog-Backend/middleware"
	r "MarvelousBlog-Backend/router"
	"MarvelousBlog-Backend/utils"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//系统是否需要重启
var RESTART_NEEDED = false

// @title PP同学个人博客接口文档
// @version 0.1
// @description MarvelousBlog-Backend Swagger接口文档
// @contact.name PP同学
// @contact.email 710955321@qq.com
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	//如果没创建博主，先创建博主账号
	CreateSuperAuthor()
	//init config之后需要重启
	if RESTART_NEEDED {
		config.Log.Infof(common.SYSTEM_INFO_LOG, "A restart is needed!")
		return
	}

	gin.SetMode(config.AppMode)
	engine := gin.New()

	//添加日志中间件
	engine.Use(middleware.LoggingMiddleware())

	//添加错误恢复中间件
	engine.Use(gin.Recovery())

	//添加跨域中间件
	//engine.Use(middleware.CORS())

	//加载各类router
	r.LoadVisitorRouters(engine)
	r.LoadAuthorRouters(engine)

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

//关闭资源
func closeResources() {
	config.Log.Infof(common.SYSTEM_INFO_LOG, "Resources Closing")
	if err := config.RedisPool.Close(); err != nil {
		config.Log.Error(common.SYSTEM_ERROR_LOG, "Redis pool closed error! errMsg = ", err)
	}
	if err := config.Db.Close(); err != nil {
		config.Log.Error(common.SYSTEM_ERROR_LOG, "Database closed error! errMsg = ", err)
	}
}

//创建博客主
func CreateSuperAuthor() {
	var superAuthor model.Author
	//检查超级作者是否已经存在
	if err := config.Db.Where("role = ?", 1).FirstOrInit(&superAuthor).Error; err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Some error happened, errMsg : ", err)
		RESTART_NEEDED = true
		return
	}
	if superAuthor.Id != 0 {
		config.Log.Infof(common.SYSTEM_INFO_LOG, "Super Author Already Exists!")
		return
	}

	//创建超级作者
	fmt.Println("Welcome! You need to initialize a Super Author!")
	fmt.Println("Please input the name of Super Author :")
	if num, err := fmt.Scanf("%s", &superAuthor.Nickname); num != 1 || err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Scan failed... \n errMsg : ", err)
		RESTART_NEEDED = true
		return
	}
	fmt.Println("Please input the password of Super Author :")
	if num, err := fmt.Scanf("%s", &superAuthor.Password); num != 1 || err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Scan failed... \n errMsg : ", err)
		RESTART_NEEDED = true
		return
	}
	superAuthor.Role = 1
	superAuthor.Avatar = common.DEFAULT_AVATAR

	//密码加密
	superAuthor.Password, _ = utils.Encrypt(superAuthor.Nickname, superAuthor.Password)

	if err := config.Db.Create(&superAuthor).Error; err != nil {
		config.Log.Errorf(common.SYSTEM_ERROR_LOG, "Some error happens while trying to create a super author! \n errMsg : ", err)
		RESTART_NEEDED = true
		return
	}
	fmt.Printf("\x1b[%d;%dmMARVELOUS! YOU HAVE CREATED A SUPER AUTHOR! \x1b[0m  \n", 47, 30)
	fmt.Println("\nNow you please restart the server!")
	RESTART_NEEDED = true
}

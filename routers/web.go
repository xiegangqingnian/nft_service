package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"io"
	"net/http"
	"nft_service/app/global/variable"
	"nft_service/app/http/controller/web"
	"nft_service/app/http/middleware/cors"
	_ "nft_service/docs"
	"os"
)

func InitRouter() *gin.Engine {
	var router *gin.Engine
	// 非调试模式（生产模式） 日志写到日志文件
	if variable.ConfigYml.GetBool("AppDebug") == false {
		//1.将日志写入日志文件
		gin.DisableConsoleColor()
		f, _ := os.Create(variable.BasePath + variable.ConfigYml.GetString("Logs.GinLogName"))
		gin.DefaultWriter = io.MultiWriter(f)
		// 2.如果是有nginx前置做代理，基本不需要gin框架记录访问日志，开启下面一行代码，屏蔽上面的三行代码，性能提升 5%
		//gin.SetMode(gin.ReleaseMode)

		router = gin.Default()
	} else {
		// 调试模式，开启 pprof 包，便于开发阶段分析程序性能
		router = gin.Default()
		//pprof.Register(router)
	}

	//根据配置进行设置跨域
	if variable.ConfigYml.GetBool("HttpServer.AllowCrossDomain") {
		router.Use(cors.Next())
	}

	url := ginSwagger.URL("http://0.0.0.0:8080/swagger/doc.json")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	//处理静态资源
	router.Static("/public", "./public")             //  定义静态资源路由与实际目录映射关系
	router.StaticFS("/dir", http.Dir("./public"))    // 将public目录内的文件列举展示
	router.StaticFile("/abcd", "./public/readme.md") // 可以根据文件名绑定需要返回的文件名

	router.POST("/generate", web.GenerateTestUsers)
	router.POST("/clean-all-user", web.CleanAllUser)

	nft := router.Group("nft/")
	{
		nft.POST("/add", web.AddNft)
		nft.POST("/search", web.GetAllNft)

	}

	//文件上传公共路由
	uploadFiles := router.Group("file/")
	{
		uploadFiles.POST("upload", web.UploadFile)
		uploadFiles.POST("up-to-ipfs", web.AddToIpfs)
		uploadFiles.POST("/download", web.Download)

	}

	return router
}

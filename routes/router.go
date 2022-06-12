package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() {
	gin.SetMode(utils.AppNode)
	// r := gin.Default()
	r := gin.New()
	r.HTMLRender = createMyRender()
	// 使用的中间件
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.Cors())

	// 加载前端静态文件
	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})

	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	/*
		后台管理路由接口
	*/
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken()) // 加入中间件，以下接口需要有权限才能使用（token正确才能调用）
	{
		// 用户模块的路由接口
		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		// 文章模块的路由接口
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)
		// 上传文件
		auth.POST("upload", v1.UpLoad)
	}

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("categorys", v1.GetCategory)
		router.GET("articles", v1.GetArticle)
		router.GET("articles/list/:id", v1.GetCateArt)
		router.GET("articles/info/:id", v1.GetSingleArticle)
		router.POST("login", v1.Login)
	}
	r.Run(utils.HttpPort)
}

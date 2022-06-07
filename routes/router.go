package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppNode)
	r := gin.Default()

	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken()) // 加入中间件，以下接口需要有权限才能使用（token正确才能调用）
	{
		// router.GET("hello", func(ctx *gin.Context) {
		// 	ctx.JSON(http.StatusOK, gin.H{
		// 		"msg": "ok",
		// 	})
		// })

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
	}
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

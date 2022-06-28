package router

import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
	e.LoadHTMLGlob("templates/*")
	e.Static("/assets", "./assets")

	e.POST("/register", controller.Register)
	e.GET("/register", controller.GoRegister)

	e.POST("/login", controller.Login)
	e.GET("/login", controller.GoLogin)

	e.GET("/blog", controller.GetBlogList)

	e.GET("/blog_details", controller.GoBlogDetails)

	e.GET("/add_blog", controller.GoAddBlogs)
	e.POST("/add_blog", controller.AddBlogs)

	e.GET("/", controller.Index)

	e.Run()
}

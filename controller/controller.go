package controller

import (
	"blog/dao"
	"blog/model"
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/russross/blackfriday/v2"
)

func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		UserName: username,
		PassWord: password,
	}
	dao.Mgr.AddUser(&user)
}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		UserName: username,
		PassWord: password,
	}
	dao.Mgr.AddUser(&user)
	c.Redirect(http.StatusFound, "/")
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := dao.Mgr.Login(username)
	fmt.Printf("user: %v\n", username)
	if user.UserName == "" {
		c.HTML(200, "login.html", "用户不存在")
		fmt.Printf("\"用户不存在\": %v\n", "用户不存在")
	} else {
		if user.PassWord != password {
			c.HTML(200, "login.html", "用户密码不正确")
		} else {
			c.HTML(200, "login.html", "登陆成功")
			c.Redirect(http.StatusFound, "/")
		}
	}
}

func GetBlogList(c *gin.Context) {
	p := dao.Mgr.GetAllBlogs()
	c.HTML(200, "blogs.html", p)
}

func GoAddBlogs(c *gin.Context) {
	c.HTML(200, "addBlog.html", nil)
}

func GoBlogDetails(c *gin.Context) {
	s := c.Query("pid")
	pid, _ := strconv.Atoi(s)
	p := dao.Mgr.GetPost(pid)
	desc_data := []byte(p.Content)
	desc_data = bytes.Replace(desc_data, []byte("\r"), nil, -1)
	content := blackfriday.Run(desc_data, blackfriday.WithNoExtensions())
	c.HTML(200, "blogDetails.html", gin.H{
		"Title":   p.Title,
		"Content": template.HTML(content),
	})

}

func AddBlogs(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}
	dao.Mgr.AddPost(&post)
	c.Redirect(http.StatusFound, "/blog")
}

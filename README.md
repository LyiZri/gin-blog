# 使用golang做一个博客系统
## 项目介绍
我们将使用golang+html+css+bootstrap+gin+gorm+mysql的技术栈模式，来构建一个前后端分离的博客系统

项目亮点
* 不再使用传统的js+html+css三大件来取写前端页面，我们将使用go来作为前端脚本语言进行构建
* 由于golang的特性，我们可以获得更高的性能，以及更快的加载速度
* 通过go我们不仅仅可以构建前端界面，也可以直接进行数据库的构建，并完成与数据库的交互


项目结构
```
    +assets //静态资源
        +css //bootstrap静态css文件
        +editor //markdown组件
        +js //bootstrap静态js文件
    +controller //控制器文件夹
        +controller.go
    +dao //连接数据库
        +dao.go
    +model //数据类型
        +model.go
    +router //路由文件
        +router.go
    +templates //html文件
        +addBlog.html //添加博客
        +blogDetails.html //博客详情
        +blogs.html //博客列表
        +header.html //页头模版
        +index.html //首页
        +login.html //登陆页面
        +register.html //注册页面
    +go.mod //mod 文件
    +go.sum //go依赖包
    +main.go //go主文件
```

项目功能介绍：
* 登陆/注册操作
* 使用markdown编辑器完成博客编写
* 进行博客发布
* 能够查看往期发布的博客，并且进行查看

## 项目依赖
项目所使用到的包

go:
> github.com/gin-gonic/gin
> gorm.io/gorm
> github.com/russross/blackfriday/v2


html:
> bootstrap
> editormd

## 项目搭建
### 项目框架搭建
首先创建以下几个文件夹以及文件
* assets 用来存放静态资源文件
* controller/controller.go 用来存放执行函数
* dao/dao.go 连接数据库文件
* model/model.go 数据库模型
* router/router.go 路由
* main.go go主文件
* templates/index.html 首页界面

执行
> go mod init main.go
生成go.mod文件

在[bootstrap官网](https://getbootstrap.com/docs/5.2/getting-started/download/)下载bootstrap包，并把js和css两个文件夹拷贝进assets文件夹下

### 实现model
首先我们已知本博客的功能，需要完成全部功能的话至少需要两张表，也就是用户表和博客表，我们进行在model中创建这两张表

首先 给这个包起个名字
```go
package model
```
然后分析用户表，
用户表的主要字段应该是用户名和密码两个点，所以我们可以这样创建
```go
type User struct {
	gorm.Model
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
```
然后引入gorm包
>   import "gorm.io/gorm"
分析博客表
博客的基本主要字段应该是标签，标题以及内容这三项
```go
type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `gorm:"type:text"`
	Tag     string `json:"tag"`
}
```
最终model文件应该是这样的
```go
package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserName string `json:"username"`
	PassWord string `json:"password"`
}
type Post struct {
	gorm.Model
	Title   string `json:"title"`
	Content string `gorm:"type:text"`
	Tag     string `json:"tag"`
}

```
### 实现dao
上面已经完成了数据库模型的构建，我们需要在dao中完成一系列关于数据库的操作

既然dao中是完成数据库操作的，那么我们首先要做的肯定是进行数据库表的创建与连接，

>首先引入dao文件
然后引入gorm建立数据库连接
引入gorm中的mysql连接包
import (
    "blog/model"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
>    )

我们首先使用mysql创建一个名为blog_db的数据库
然后建立数据库连接
```go
    //charset=utf8为字符串解析
	//parseTime=true不添加的话 查询时会报错，Scan error on column index 1, name "created_at": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
dsn := "root:your_password@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```
root后跟的是你自己的数据库密码

创建user以及post表
```go
Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
```
接下来只要引入dao包就可以正常的使用数据库了


### 实现首页的编写
我们可以在bootstrap上找到一个header的模版并直接进行套用
我使用的模版是直接在[这里](https://getbootstrap.com/docs/5.2/examples/headers/)找的

![16563963604474.jpg](https://p9-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/a7e3ec9f5bf74d67b082c99d9d92b97b~tplv-k3u1fbpfcp-watermark.image?)

当然，拿到模版后我们需要进行一些小小的修改
首先进行模块化编写
在header.html文件的开头写上
>{{define "header"}}

结尾处编写
> {{end}}

然后在中间内容中写入
```html
<header class="p-3 bg-dark text-white">
    <div class="container">
      <div class="d-flex flex-wrap align-items-center justify-content-center justify-content-lg-start">
        <a href="#" class="navbar-brand">lyi</a>

        <ul class="nav col-12 col-lg-auto me-lg-auto mb-2 justify-content-center mb-md-0">
          <li><a href="/" class="nav-link px-2 text-secondary">首页</a></li>
          <li><a href="/blog" class="nav-link px-2 text-white">博客</a></li>
          <li><a href="/add_blog" class="nav-link px-2 text-white">添加博客</a></li>
        </ul>
        <div class="text-end">
          <a type="button" href="/login" class="btn btn-outline-light me-2">登陆</a>
          <a type="button" href="/register" class="btn btn-warning">注册</a>
        </div>
      </div>
    </div>
  </header>
```
这其中herf的标签已经被我写好了，用于后续的路由跳转

**这样，一个header的template就写好了**

接下来在index.html中引入并使用header
在头部直接引入bootstrap文件
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/assets/css/bootstrap.min.css">
    <script src="assets/js/bootstrap.min.js"></script>
</head>

<body>
    <div class="container">
        
        {{template "header"}}
        首页...
</div>

</body>

</html>
```
我们就得到了这样的页面


### 实现controller
为了能够看到index.html的真容，我们需要对index.html文件进行载入

```go

    func Index(c *gin.Context){
        c.HTML(200,"index.html",nil)
    }
```
函数的形参c的类型为*gin.Context这是gin中一个非常经典的上下文控制类型，我们可以使用这个类型中的一些方法来完成对html的控制，gin是构建在golang http库之上的，他的功能和http库很接近，但是会更容易于使用

c.HTML参数解析：
* 200为http成功响应状态码
* index.html代表我们需要链接的文件
* nil 表示所需传入的数据为空

### 实现router
引入controller包和gin包
```go
import (
	"blog/controller"

	"github.com/gin-gonic/gin"
)
```

定义一个start方法，
创建一个gin实例
```go
e := gin.Default()
```
获取来自templates下的文件
```go
e.LoadHTMLGlob("template/*")
```
加载静态文件
```go
e.Static("/assets","./assets")
```

讲index.html文件显示在跟路由下
```go
e.GET("/",controller.Index)
```
gin实例运行
```go
e.Run()
```
打开网页localhost:8080就可以看到我们的界面

![16563966112986.jpg](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/a4f64a847e7342888282b6f48f099d14~tplv-k3u1fbpfcp-watermark.image?)
### 实现登陆注册
上面已经实现过用户表了，所以注册就是往用户表中添加数据，登陆就是从用户表中查询数据并进行验证

注册添加用户实现（controller.go）
```go
//获取表单输入数据
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
//用户添加操作
func AddUser(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")

	user := model.User{
		UserName: username,
		PassWord: password,
	}
	dao.Mgr.AddUser(&user)
}
```
dao.go
```go
//向数据库中添加新注册用户数据
func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(user)
}
```

登陆查询用户实现
```go
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
```
dao.go
```go
//查询用户名为username的数据
func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("user_name=?", username).First(&user)
	fmt.Printf("user: %v,psw:%v\n", user.UserName, user.PassWord)
	return user
}
```

### 实现注册和登录界面
注册：
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/assets/css/bootstrap.min.css">
    <script src="assets/js/bootstrap.min.js"></script>
    <title>注册</title>
</head>

<body>
    <div class="container">
        
        {{template "header"}}
        
        <!--代表使用post方法-->
        <form method="post" action="/register">
            <div class="mb-3">
              <label for="exampleInputEmail1" class="form-label">用户名</label>
              <!--name为gin取数据依赖项-->
              <input type="text" name="username" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp">
            </div>
            <div class="mb-3">
              <label for="exampleInputPassword1" class="form-label">Password</label>
              <input type="password" name="password" class="form-control" id="exampleInputPassword1">
            </div>
            <div class="mb-3">
              <label for="exampleInputPassword1" class="form-label">Password</label>
              <input type="password" name="password2" class="form-control" id="exampleInputPassword1">
            </div>
            <button type="submit" class="btn btn-primary">注册</button>
          </form>
    </div>

</body>

</html>
```
contraller中register页面注册
```go
func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}
```


登陆
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/assets/css/bootstrap.min.css">
    <script src="assets/js/bootstrap.min.js"></script>
    <title>登陆</title>
</head>

<body>
    <div class="container">
        
        {{template "header"}}
        <form method="post" action="/login">
          <p style="background-color: red;">{{.}}</p>
            <div class="mb-3">
              <label for="exampleInputEmail1" class="form-label">用户名</label>
              <input type="text" name="username" class="form-control" id="exampleInputEmail1" aria-describedby="emailHelp">
            </div>
            <div class="mb-3">
              <label for="exampleInputPassword1" class="form-label">Password</label>
              <input type="password" name="password" class="form-control" id="exampleInputPassword1">
            </div>
            <button type="submit" class="btn btn-primary"> 登陆</button>
          </form>
    </div>

</body>

</html>



```
登陆界面注册
```go
func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}
```

router内嵌入路由，并导入方法
```go
	e.POST("/register", controller.Register)
	e.GET("/register", controller.GoRegister)

	e.POST("/login", controller.Login)
	e.GET("/login", controller.GoLogin)
```

因为我们在前面已经在header中创建过了登陆和注册的入口，所以打开我们就可以看到页面，并且进行登录注册的操作

![16563996548206.jpg](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/9746c37572c846ab9142f386d11dd877~tplv-k3u1fbpfcp-watermark.image?)

我们访问数据库，使用
```sql
select * from users;
```
也可以查询到我们刚刚创建的用户数据


### 集成markdown编写器
在这里我们使用到的markdown编写器，名为[editormd](https://pandao.github.io/editor.md/en.html),可以访问官网并进行下载

将下载下来的文件放入assets中，因为原文件名太长，我们在这里就直接更名为editor

在addBlog.html文件中使用editor
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="/assets/css/bootstrap.min.css">
    <script src="assets/js/bootstrap.min.js"></script>
    <link rel="stylesheet" href="/assets/editor/css/editormd.min.css">
    <title>博客编写</title>
</head>

<body>
    <div class="container">
        {{template "header"}}
        <form action="/add_blog" method="post">
        <div class="row">
            <div class="col-8">
                <div id="test-editor">
                    <textarea style="display:none;" name="content">
                    </textarea>
                </div>
                <script src="https://cdn.bootcss.com/jquery/1.11.3/jquery.min.js"></script>
                <!-- src路径需要进行修改 -->
                <script src="/assets/editor/editormd.amd.min.js"></script>
                <script type="text/javascript">
                    $(function () {
                        var editor = editormd("test-editor", {
                            // width  : "100%",
                            height: 640,
                            path: "/assets/editor/lib/"
                        });
                    });
                </script>
            </div>
            <div class="col-4">

                    <div class="">
                        <label for="exampleFormControlInput1" class="form-label">标题</label>
                        <input type="text" class="form-control" id="exampleFormControlInput1" name="title">
                    </div>
                    <div class="">
                        <label for="exampleFormControlInput1" class="form-label">标签</label>
                        <input type="text" class="form-control" id="exampleFormControlInput1" name="tag">
                    </div>
                    <button type="submit" class="btn btn-primary" method="post">发布</button>
                    
                </div>
            </div>
        </form>
        </div>
        
    </body>
    
    </html>
```
我们将editormd编辑器放置在了左侧，而在右侧放了表单，进行博客的提交
这里有几个点需要注意
```js
$(function () {
                        var editor = editormd("test-editor", {
                            // width  : "100%",
                            height: 640,
                            path: "/assets/editor/lib/"
                        });
                    });

```
第一个参数为我们想要显示的地方的id，path代表我们将editor文件存放在了哪个位置，
>   form标签一定要置于最外层，包裹着markdown编写器和title以及tag输入框，否则将获取不到数据

#### 将信息存入数据库
```go
//添加博客信息
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

//注册页面
func GoAddBlogs(c *gin.Context) {
	c.HTML(200, "addBlog.html", nil)
}
```

dao文件中添加数据
```go
func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}
```
>   c.Redirect(http.StatusFound, "/blog")代表重定向至/blog


router添加路由，注册方法
```go
	e.GET("/add_blog", controller.GoAddBlogs)
	e.POST("/add_blog", controller.AddBlogs)
```

![16564004169382.jpg](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/48795e5deb664335bd8b01e99b95c4bc~tplv-k3u1fbpfcp-watermark.image?)
### 博客列表

博客列表页面实现
```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>博客</title>
    <link rel="stylesheet" href="/assets/css/bootstrap.min.css">
    <script src="/assets/js/bootstrap.min.js"></script>
</head>

<body>
    <div class="container">
        {{template "header"}}
        <div class="row">
            {{range $post := . -}}
            <div class="col-md-6">
                <div
                    class="row g-0 border rounded overflow-hidden flex-md-row mb-4 shadow-sm h-md-250 position-relative">
                    <div class="col p-4 d-flex flex-column position-static">
                        <strong class="d-inline-block mb-2 text-primary">分类</strong>
                        <h3 class="mb-0">{{$post.Title}}</h3>
                        <p class="card-text mb-auto">{{$post.Content}}</p>
                        <a href="/blog_details?pid={{$post.ID}}" class="stretched-link">查看详情</a>
                    </div>
                    <div class="col-auto d-none d-lg-block">
                        <svg class="bd-placeholder-img" width="200" height="250" xmlns="http://www.w3.org/2000/svg"
                            role="img" aria-label="Placeholder: Thumbnail" preserveAspectRatio="xMidYMid slice"
                            focusable="false">
                            <title>Placeholder</title>
                            <rect width="100%" height="100%" fill="#55595c"></rect><text x="50%" y="50%" fill="#eceeef"
                                dy=".3em">Thumbnail</text>
                        </svg>

                    </div>
                </div>
            </div>
            {{end}}
        </div>
    </div>
</body>

</html>
```

>   {{range $post := . -}}代表将传入的切片进行循环，每个被循环的元素名为 `$`post

dao获取数据
```go
func (mgr *manager) GetAllBlogs() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}
```

controller解析数据，注册页面
```go
func GetBlogList(c *gin.Context) {
	p := dao.Mgr.GetAllBlogs()
	c.HTML(200, "blogs.html", p)
}
func GoAddBlogs(c *gin.Context) {
	c.HTML(200, "addBlog.html", nil)
}
```

router注册页面，注册方法
```go
	e.GET("/blog", controller.GetBlogList)
```

### 获取博客详情
因为我们上传到数据库中的markdown文件是一个字符串类型的，所以如果我们想要以文档形式转换过来，就必须使用到一点点`小技术` 在这里我们使用到了`blackfriday`包

```go

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
```
先把我们得到的markdown字符串转为byte数组，再经过转译，最后使用go内置的template包，来传给html文件

博客页面的html文件编写很简单，看起来就好像是只需要进行接收数据并展示就可以了
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>博客详情</title>
    <link rel="stylesheet" href="/assets/css/bootstrap.min.css">
    <script src="/assets/js/bootstrap.min.js"></script>
</head>
<body>
    <div class="container">
        {{template "header"}}
        <div class="row">
            <div class="col-md-12">
                <h1>{{.Title}}</h1>
                <p>
                    {{.Content}}
                </p>
            </div>
        </div>
    </div>
</body>
</html>
```

![16564013725520.jpg](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/5abe39044ba64b1a90a15c397ebd55ac~tplv-k3u1fbpfcp-watermark.image?)

![16564014222204.jpg](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/d5f19769411a48db9a060fdfb57300d9~tplv-k3u1fbpfcp-watermark.image?)






# gin-blog

package dao

import (
	"blog/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Manager interface {
	AddUser(user *model.User)
	Login(username string) model.User
	AddPost(post *model.Post)
	GetAllBlogs() []model.Post
	GetPost(pid int) model.Post
}

type manager struct {
	db *gorm.DB
}

func (mgr *manager) AddUser(user *model.User) {
	mgr.db.Create(user)
}
func (mgr *manager) AddPost(post *model.Post) {
	mgr.db.Create(post)
}

var Mgr Manager

func init() {
	//charset=utf8为字符串解析
	//parseTime=true不添加的话 查询时会报错，Scan error on column index 1, name "created_at": unsupported Scan, storing driver.Value type []uint8 into type *time.Time
	dsn := "root:1296729980@tcp(127.0.0.1:3306)/go_db?charset=utf8&parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	Mgr = &manager{db: db}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Post{})
}

func (mgr *manager) Login(username string) model.User {
	var user model.User
	mgr.db.Where("user_name=?", username).First(&user)
	fmt.Printf("user: %v,psw:%v\n", user.UserName, user.PassWord)
	return user
}
func (mgr *manager) GetAllBlogs() []model.Post {
	var posts = make([]model.Post, 10)
	mgr.db.Find(&posts)
	return posts
}
func (mgr *manager) GetPost(pid int) model.Post {
	var post model.Post
	mgr.db.First(&post, pid)
	return post
}

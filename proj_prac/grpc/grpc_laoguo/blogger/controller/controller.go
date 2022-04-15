package controller

import (
	"blog/dao"
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user := model.User{
		Username: username,
		Password: password,
	}
	dao.Mgr.Register(&user)
	c.Redirect(301, "/")
}
func GoRegister(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func GoLogin(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	u := dao.Mgr.Login(username)
	if u.Username == "" {
		c.HTML(200, "login.html", "user does not exist")
		fmt.Println("User does not exist")
	} else {
		if u.Password != password {
			fmt.Println("password is wrong")
			c.HTML(200, "login.html", "password is wrong")

		} else {
			fmt.Println("Login successfully")
			c.Redirect(301, "/")
		}
	}

}

func Index(c *gin.Context) {
	c.HTML(200, "index.html", nil)

}

func ListUser(c *gin.Context) {
	c.HTML(200, "userlist.html", nil)

}

func GetPostIndex(c *gin.Context) {
	posts := dao.Mgr.GetAllPost()
	c.HTML(200, "postIndex.html", posts)
}

func AddPost(c *gin.Context) {
	title := c.PostForm("title")
	tag := c.PostForm("tag")
	content := c.PostForm("content")

	post := model.Post{
		Title:   title,
		Tag:     tag,
		Content: content,
	}
	dao.Mgr.AddPost(&post)
	c.Redirect(302, "/post_index")
}

func GoAddPost(c *gin.Context) {
	c.HTML(200, "post.html", nil)

}

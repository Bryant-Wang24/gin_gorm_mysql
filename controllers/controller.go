package controllers

import (
	"gin/models"

	"github.com/gin-gonic/gin"
)

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	user, _ := models.GetUser(username)
	if user.ID == 0 {
		c.JSON(200, gin.H{
			"code":    1,
			"message": "用户名不存在",
		})
	} else {
		if user.Password == password {
			c.JSON(200, gin.H{
				"code":    0,
				"message": "登录成功",
				"data":    user,
			})

		} else {
			c.JSON(200, gin.H{
				"code":    1,
				"message": "用户名或密码错误",
			})
		}
	}
}

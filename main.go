package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type User struct {
	gorm.Model
	Username string
	Password string
}

// 创建表
func create(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

// 插入数据
func insert(db *gorm.DB) {
	db.Create(&Product{Code: "L1212", Price: 100})
}

// 查询
func find(db *gorm.DB) {
	var p Product
	db.First(&p, 1) // 根据主键查询： 查询id为1的product
	fmt.Printf("p: %v\n", p)
	db.First(&p, "code = ?", "L1212") // 根据条件查询： 查询code字段为1001的product
	result := db.Find(&p)
	fmt.Println("result", result.RowsAffected)
	fmt.Printf("p: %v\n", p)
}

// 更新
func update(db *gorm.DB) {
	var p Product
	// db.First(&p, 3)``
	db.First(&p, "price=?", 500)
	// db.Model(&p).Updates(Product{Code: "1005", Price: 500})
	db.Model(&p).Updates(map[string]interface{}{"price": 900, "code": "1009"})
	// p.Price = 200
	// db.Save(&p)
}

// 删除
func delete(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	db.Delete(&p) // 软删除，并不会真的删除数据，会在表里添加一个标记字段
}

func main() {
	dsn := "root:485969746wqs@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/admin/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		fmt.Println(username, password)
		// 查询用户是否存在
		var user User
		db.Where("username = ?", username).First(&user)
		fmt.Println("user", user.ID)
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
	})
	// insert(db)
	// find(db)
	// update(db)
	// delete(db)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}

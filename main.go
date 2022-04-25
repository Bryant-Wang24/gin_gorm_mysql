package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
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
	// insert(db)
	// find(db)
	// update(db)
	delete(db)
}

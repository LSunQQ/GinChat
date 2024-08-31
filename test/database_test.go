package test

import (
	"ginchat/models"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestDataBase(t *testing.T) {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(172.24.87.200:3306)/ginchat?charset=utf8mb4&parseTime=true&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{
		Name:     "test",
		PassWord: "123",
		Email:    "123",
	}
	db.Create(user)
	// db.Create(&Product{Code: "D42", Price: 100})

	// Read
	// var product Product
	db.First(&user, 1) // 根据整型主键查找
	// db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&user).Update("PassWord", "654321")
	// Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	// db.Delete(&product, 1)
}

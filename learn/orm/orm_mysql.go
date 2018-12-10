package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/go-sql-driver/mysql"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type ProductMysql1 struct {
	gorm.Model
	Code  string ``
	Price uint
}

func main() {
	db, err := gorm.Open("mysql", "dev:Aa111111@(112.74.185.158:3306)/dandantuan?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("连接数据库失败")
	}
	defer db.Close()

	// 自动迁移模式
	db.AutoMigrate(&ProductMysql1{})

	// 创建
	db.Create(&ProductMysql1{Code: "L1212", Price: 1000})

	// 读取
	var product ProductMysql1
	db.First(&product, 1)                   // 查询id为1的product
	db.First(&product, "code = ?", "L1212") // 查询code为l1212的product

	// 更新 - 更新product的price为2000
	db.Model(&product).Update("Price", 2000)

	db.Save(&product)

	// 删除 - 删除product
	//db.Delete(&product)
}

// 设置User的表名为`profiles`
func (ProductMysql1) TableName() string {
	return "profiles"
}
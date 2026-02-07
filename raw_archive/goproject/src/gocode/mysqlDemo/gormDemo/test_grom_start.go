package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code string
	Pric uint
}

//创建表
func create(db *gorm.DB) {
	db.AutoMigrate(&Product{})
}

//插入数据
func insert(db *gorm.DB) {
	p := Product{
		Code: "1001",
		Pric: 100,
	}

	db.Create(&p)
}

//查询
func find(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	fmt.Printf("p: %v\n", p)
	db.First(&p, "code = ?", "1001") //查找code字段为1001的记录
	fmt.Printf("p: %v\n", p)
}

//更新
func update(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	db.Model(&p).Update("Pric", 1000)
	db.Model(&p).Updates(Product{Pric: 1001, Code: "1002"}) //仅更新非零值字段
	db.Model(&p).Updates(map[string]interface{}{"Pric": 1003, "Code": "1003"})
}

func del(db *gorm.DB) {
	var p Product
	db.First(&p, 1)
	db.Delete(&p, 1)
}

func main() {
	dsn := "root:chenjie110@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gorm err")
	}

	//创建表(把结构体直接创建到mysql)
	//

	//根据主键来查询
	//find(db)
	update(db)
	del(db)
}

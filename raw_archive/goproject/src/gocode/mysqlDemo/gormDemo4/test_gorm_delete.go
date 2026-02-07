package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dsn := "root:chenjie110@tcp(127.0.0.1:3306)/golang_db?charset=utf8mb4&parseTime=True"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("gorm err")
	}

	db = d
}

type User struct {
	gorm.Model
	Name     string
	Age      int
	Birthday time.Time
}

func creatTable() {
	db.AutoMigrate(&User{}) //创建一个表的结构
}

func del1() {
	var user User
	db.First(&user)
	//db.Delete(&user) // soft delete
}

func del2() {
	db.Delete(&User{}, 11)
}

func del3() {
	db.Delete(&User{}, []int{1, 2, 3, 4, 5})
}

/*
func (user *User) delete() {
	db.Unscoped().Delete(&user)
}*/

//hook
func (u *User) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println("BeforeDelete...")
	return
}

func insert() {
	user := User{
		Name:     "tom",
		Age:      20,
		Birthday: time.Now(),
	}
	db.Create(&user)
}

func main() {

	// db.Where("name = ?","jinzhu").Delete(&user)
	//insert()
	//del2()
}

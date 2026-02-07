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

var user = User{
	Name:     "tom",
	Age:      20,
	Birthday: time.Now(),
}

func insert1() { //创建一条记录

	result := db.Create(&user) //调用Create()来创建表

	fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected) //result.RowsAffected表示影响了几行
	fmt.Printf("user.ID:%v\n", user.ID)
}

func insert2() {
	db.Select("Name", "Age", "CreatedAt").Create(&user) //Select:选择哪些元素需要添加
}

func insert3() {
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	db.Create(&users)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	fmt.Println("BeforeCreate..")
	return
}

func insert4() {
	db.Model(&User{}).Create(map[string]interface{}{
		"Name": "jinzhu", "Age": 18,
	})
}

func main() {
	//creatTable() //调用的时候表就已经创建完毕了
	//insert1()
	//insert2()
	//insert3()
	insert4()
}

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

func select1() {
	var user User
	/*
		db.First(&user) //获取第一条记录
		//fmt.Printf("user: %v\n", user)
		fmt.Printf("user.ID:%v\n", user.ID)
		db.Take(&user) //获取一条记录
		fmt.Printf("user.ID: %v\n", user.ID)
	*/
	db.Last(&user) //获取最后一条记录
	fmt.Printf("user.ID: %v\n", user.ID)
}

func select2() {
	//var user User
	//db.First(&user, 10)//根据id来查询
	//fmt.Printf("user.ID: %v\n", user.ID)
	//SELECT * FROM users WHERE id = 10;

	//db.First(&user, "10") //查询id(字符串s)
	//fmt.Printf("user.ID: %v\n", user.ID)
	//SELECT * FROM users WHERE id = 10;

	var users []User
	db.Find(&users, []int{1, 2, 3}) //Find(查找全部)
	for _, user := range users {
		fmt.Printf("user.ID: %v\n", user.ID)
	}
	//SELECT * FORM users WHERE id IN (1,2,3);

}

func select3() { //一共查了几条记录
	var users []User
	result := db.Find(&users)
	fmt.Printf("result.RowsAffected: %v\n", result.RowsAffected)
}

func main() {
	//creatTable() //调用的时候表就已经创建完毕了
	//select1()
	select3()
}

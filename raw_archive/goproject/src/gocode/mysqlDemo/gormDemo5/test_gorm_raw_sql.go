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

func testRaw1() { //原生数据库
	type Result struct {
		ID   int
		Name string
		Age  int
	}

	var result Result
	db.Raw("select id,name,age from users where name = ?", "tom").Scan(&result)
	fmt.Printf("result: %v\n", result)

	var age int
	db.Raw("select sum(age) from users").Scan(&age)
	fmt.Printf("age: %v\n", age)
}

func main() {
	//testRaw1()

	user := User{
		Name:     "kite",
		Age:      21,
		Birthday: time.Now(),
	}
	db.Create(&user)

}

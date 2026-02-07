/*
package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "root:chenjie110@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
		db, err := sql.Open("mysql", "root:chenjie110@/go_db")
		if err != nil {
			panic(err)
		}
		fmt.Print(db)
		// See "Important settings" section.
		//最大连接时间:3分钟
		db.SetConnMaxLifetime(time.Minute * 3)
		//最大连接数
		db.SetMaxOpenConns(10)
		//空闲连接数
		db.SetMaxIdleConns(10)
		fmt.Printf("db:%v\n", db)
	err := initDB()
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Println("连接成功")
	}
}
*/
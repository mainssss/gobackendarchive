package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //DB:Database(数据库)

func initDB() (err error) {
	dsn := "root:chenjie110@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4&parseTime=True"
	db, err = sql.Open("mysql", dsn) //open只是检验代码正确与否，并不会创建连接，连接需要ping
	if err != nil {
		return err
	}
	err = db.Ping() //只有ping成功了，才算是真正连接上了！！！
	if err != nil {
		return err
	}
	return nil
}

func insert() { //输入数据表
	//插入、更新、删除操作一般使用Exec()方法
	s := "insert into user_tbl (username,password) values(?,?)"
	r, err := db.Exec(s, "zhangsan", "zs123")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		i, _ := r.LastInsertId() //查询最后一次操作后插入数据的ID
		fmt.Printf("i:%v\n", i)
	}
}

type User struct {
	id       int
	username string
	password string
}

func queryOneRow() {
	s := "select * from user_tbl where id = ?" //查询user—_tbl表中id =？的数据
	var u User
	//QueryRow进行一次查询，最多返回一行结果，QueryRow总是返回非nil的值，直到被Scan()调用，才会返回被延时的错误
	err := db.QueryRow(s, 2).Scan(&u.id, &u.username, &u.password) //Scan:将查询到的数据赋值给User结构体
	if err != nil {
		fmt.Printf("err : %v\n", err)
	} else {
		fmt.Printf("u:%v\n", u) //打印结构体
	}

}

func queryManyRow() {
	//"select * from user_tbl id>0"    查询id>0的条件
	s := "select * from user_tbl" //sql命令,功能:查询整个数据表
	r, err := db.Query(s)         //Query()：一次返回多条记录
	//r:每行
	var u User
	defer r.Close() //关闭r释放数据库连接!!!
	if err != nil {
		fmt.Printf("err:%v", err)
	} else {
		//循环读取结果集中的数据
		for r.Next() {
			r.Scan(&u.id, &u.username, &u.password)
			fmt.Printf("u:%v", u)
		}
	}
}

func main() {

	err := initDB() //连接数据库函数，返回一个err
	if err != nil {
		fmt.Printf("err:%v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	queryManyRow()
}

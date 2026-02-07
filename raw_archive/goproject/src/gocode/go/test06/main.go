package main

//连接mysql实例、
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //()init
)

func query() {

}

func insert() {

}

func main() {
	dsn := "root:chenjie110@tcp(127.0.0.1:3306)/goday10"
	//连接数据库信息
	db, err := sql.Open("mysql", dsn) //open不会校验用户名准不准确，只会校验dsn格式对不对
	if err != nil {
		fmt.Printf("OPEN %s failed,err:%v\n", dsn, err)
		return
	}
	err = db.Ping() //校验dsn是否正确
	if err != nil {
		fmt.Printf("err :%v\n", err)
		return
	}
	fmt.Println("数据库连接成功...")

}

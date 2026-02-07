package main

//连接mysql实例、
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //()init
)

var db *sql.DB //db是一个连接池对象

func initDB() (err error) {
	dsn := "root:chenjie110@tcp(127.0.0.1:3306)/sql_test"
	//连接数据库信息
	db, err = sql.Open("mysql", dsn) //open不会校验用户名准不准确，只会校验dsn格式对不对
	if err != nil {
		return
	}
	err = db.Ping() //校验dsn是否正确
	if err != nil {
		return
	}

	db.SetMaxOpenConns(10) //设置数据库的最大连接数
	db.SetMaxIdleConns(5)  //设置最大空闲连接数
	return
}

type user struct {
	id   int
	name string
	age  int
}

//查询
func query(id int) {
	var u1 user
	//1.查询单条记录的sql语句
	sqlStr := `select id,name,age from user where id=?;`
	//2、执行
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.name, &u1.age) //从连接池取出一条连接取数据库查询单条记录,赋值
	//必须对rowObj使用scan方法，它会释放连接
	//3.拿到结果

	fmt.Printf("u1:%#v\n", u1)
}

func insert() {

}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println("连接数据库成功...")
	query(2)
}

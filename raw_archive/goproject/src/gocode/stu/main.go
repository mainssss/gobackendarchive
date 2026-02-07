package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB //db是一个连接池对象
var num int
var name string
var grade float32
var xz int

//var id int

type student struct {
	num   int
	name  string
	grade float32
}

func initDB() (err error) {
	dsn := "root:chenjie110@tcp(127.0.0.1:3306)/go_stu"
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

func xiugaiXue1() {
	fmt.Println("请输入修改后的正确学号:")
	fmt.Scanf("%d\n", &num)
	fmt.Println("请输入该学生的姓名:")
	fmt.Scanf("%s\n", &name)
	xiugaiXue(num, name)
}

func xiugaiXue(num int, name string) {
	sqlStr := `update student set num=? where name=?` //改成绩,用名字查询/改
	ret, err := db.Exec(sqlStr, num, name)
	if err != nil {
		fmt.Println("修改失败:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("您更新了%d行", n)
}
func xiugaiName1() {
	fmt.Println("请输入您更改后的名字")
	fmt.Scanf("%s\n", &name)
	fmt.Println("请输入要此学生学号")
	fmt.Scanf("%d\n", &num)
	xiugaiName(name, num)
}

func xiugaiName(name string, num int) {
	sqlStr := `update student set name=? where num=?` //改成绩,用名字查询/改
	ret, err := db.Exec(sqlStr, name, num)
	if err != nil {
		fmt.Println("修改失败:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("您更新了%d行", n)
}

func xiugai1() {
	fmt.Println("请输入您要修改的成绩")
	fmt.Scanf("%f\n", &grade)
	fmt.Println("请输入此学生姓名")
	fmt.Scanf("%s\n", &name)
	xiugai(grade, name)
}

func xiugai(grade float32, name string) {
	sqlStr := `update student set grade=? where name=?` //改成绩,用名字查询/改
	ret, err := db.Exec(sqlStr, grade, name)
	if err != nil {
		fmt.Println("修改失败:", err)
		return
	}
	n, err := ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("您更新了%d行", n)
}

func chaxun(xz int) {
	var u1 student
	//1.查询单条记录的sql语句
	sqlStr := `select num,name,grade from student where num=?;`
	//2、执行
	db.QueryRow(sqlStr, xz).Scan(&u1.num, &u1.name, &u1.grade) //从连接池取出一条连接取数据库查询单条记录,赋值
	//必须对rowObj使用scan方法，它会释放连接
	//3.拿到结果
	fmt.Println("学生的学号为:", u1.num)
	fmt.Println("学生的姓名为:", u1.name)
	fmt.Println("学生的成绩为:", u1.grade)
	//fmt.Printf("u1:%#v\n", u1)
}

func charu(num int, name string, grade float32) {
	//sql语句
	sqlStr := `insert into student(num,name,grade) values(?,?,?)`
	//exec
	ret, err := db.Exec(sqlStr, num, name, grade)
	if err != nil {
		fmt.Println("插入失败:", err)
		return
	}
	theId, err := ret.LastInsertId() //插入数据的id
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("你插入数据的id为:%d\n", theId)
}

func xianshi() {
	sqlStr := `select num,name,grade from student where id>?`
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Println("抱歉,查询出现错误")
		return
	}
	defer rows.Close() //释放数据库连接
	for rows.Next() {
		var u student
		err := rows.Scan(&u.num, &u.name, &u.grade)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("学号:%d\n姓名:%s\n成绩:%g\n\n", u.num, u.name, u.grade)
	}
}

func Del1() {
	fmt.Println("请输入您要删除的学生学号")
	fmt.Scanf("%d\n", &num)
	Del(num)
}

func Del(num int) {
	sqlStr := `delete from student where num=?`
	ret, err := db.Exec(sqlStr, num)
	if err != nil {
		fmt.Println("删除出现错误:", err)
		return
	}
	n, err := ret.RowsAffected() //影响几行
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("您修改了%d行", n)
}

func main() {

	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Println("连接数据库成功...")
	var n int
	for {
		fmt.Println(`
	----------欢迎使用学生管理系统----------
	1.添加学生
	2.查询学生
	3.修改
	4.显示所有学生
	5.删除学生
	6.退出`)
		fmt.Println("请输入你的选择")

		fmt.Scanf("%d\n", &n)
		switch n {
		case 1:
			fmt.Println("添加学生")
			fmt.Println("请输入学生id:")
			fmt.Scanf("%d\n", &num)
			fmt.Println("请输入学生名字:")
			fmt.Scanf("%s\n", &name)
			fmt.Println("请输入学生成绩:")
			fmt.Scanf("%f\n", &grade)
			charu(num, name, grade)
		case 2:
			fmt.Println("查询学生")
			fmt.Println("请输入学号:")
			fmt.Scanf("%d\n", &xz)
			chaxun(xz)
		case 3:
			fmt.Println("修改")
			fmt.Println(`
			1.修改学号
			2.修改名字
			3.修改成绩
			4.返回上级`)
			fmt.Println("请输入你的选择:")
			fmt.Scanf("%d\n", &xz)
			switch xz {
			case 1:
				xiugaiXue1()
			case 2:
				xiugaiName1()
			case 3:
				xiugai1()
			case 4:
				break
			default:
				fmt.Println("您的输入有误,请重新输入")
			}
			//xiugai1()
		case 4:
			xianshi()
		case 5:
			Del1()
		case 6:
			fmt.Println("你确定要退出吗?")
			fmt.Println(`1.确定     2.否`)
			fmt.Scanf("%d\n", &xz)
			switch xz {
			case 1:
				os.Exit(0)
			case 2:
				break
			default:
				fmt.Println("你的输入有误,请重新输入...")
			}
		default:
			fmt.Println("您的输入有误,请重新输入")
		}
	}

}

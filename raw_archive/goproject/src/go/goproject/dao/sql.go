package dao

import (
	"Aurora/struct/studentExercise/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// DB 定义一个全局变量 DB
var DB *gorm.DB

func InitMysql() (err error) {
	dsn := "root:zjl@tcp(127.0.0.1:3306)/student?charset=utf8mb4&parseTime=true&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
         //fmt.Println(err)
         return err
	}
	return nil

}


func Query(order string , s string){
	rows , err := DB.DB().Query(order,s)
	if err != nil{
		panic(err)
	}
	for rows.Next(){
		var stu model.Student
		rows.Scan(&stu.Id,&stu.Name,&stu.Class,&stu.Score)
		fmt.Printf("id : %d  name : %s  class : %s  score : %d\n",stu.Id , stu.Name , stu.Class , stu.Score)
	}
}

// AfterUpdate 输出更改之后学生的信息
func AfterUpdate(n interface{}){
	var stu model.Student
	DB.First(&stu,n)
	fmt.Println("该学生更改之后的信息：")
	fmt.Printf("id : %d  name : %s  class : %s  score : %d\n",stu.Id , stu.Name , stu.Class , stu.Score)
}

func Update(n interface{} , s string , id int){
	stmt , err := DB.DB().Prepare(s)
	if err != nil{
		panic(err)
	}
	stmt.Exec(n , id)
}
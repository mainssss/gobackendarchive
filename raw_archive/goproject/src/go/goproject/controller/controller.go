ackage controller

import (
	"Aurora/struct/studentExercise/dao"

	"Aurora/struct/studentExercise/model"
	"fmt"
)

// Add 添加学生
func Add() {
	var id , score int
	var name , class string
	fmt.Println("请输入学生的id：")
	fmt.Scanln(&id)
	var stu model.Student
	dao.DB.First(&stu,id)
	if id == stu.Id{
		fmt.Println("id已存在！！！")
		return
	}
	fmt.Println("请输入学生的name：")
	fmt.Scanln(&name)
	fmt.Println("请输入学生的class：")
	fmt.Scanln(&class)
	fmt.Println("请输入学生的score：")
	fmt.Scanln(&score)
	var student = &model.Student{
		Id: id,
		Name: name,
		Class: class,
		Score: score,
	}
	err :=dao.DB.Create(&student)
	if err != nil{
		//fmt.Println(err)
		return
	}
}

// ShowList 展示学生列表
func ShowList(){
	rows , err := dao.DB.DB().Query("SELECT * FROM students ")
	if err != nil{
		panic(err)
	}
	for rows.Next(){
		var stu model.Student
		rows.Scan(&stu.Id,&stu.Name,&stu.Class,&stu.Score)
		fmt.Printf("id : %d  name : %s  class : %s  score : %d\n",stu.Id , stu.Name , stu.Class , stu.Score)
	}
}

func ShowMenu(){
	fmt.Println("-------学生管理系统-------")
	fmt.Println("-------1.添加学生 -------")
	fmt.Println("-------2.查询学生信息 ----")
	fmt.Println("-------3.编辑学生信息-----")
	fmt.Println("-------4.删除学生 -------")
	fmt.Println("-------5.展示学生列表-----")
	fmt.Println("-------6.退出系统 -------")
}

// Delete 删除学生
func Delete(){
	var id int
	fmt.Println("请输入你要删除学生的id ：")
	fmt.Scanln(&id)
	dao.DB.Where("id = ?",id).Delete(&model.Student{})
}


func FindList(){
	fmt.Println("--------查询学生信息-------")
	fmt.Println("-------1.使用学号查询------")
	fmt.Println("-------2.使用姓名查询------")
	fmt.Println("-------3.使用班级查询------")
	fmt.Println("-------4.使用成绩查询------")
	fmt.Println("-------5.退出查询---------")

}

// Find 按照条件查找
func Find() {
	var number int
	for {
		FindList()
		fmt.Println("请输入你的选择：")
		fmt.Scanln(&number)
		switch number {
		case 1:
			fmt.Println("使用学号查询")
			var id int
			fmt.Println("请输入要查询学生的学号：")
			fmt.Scanln(&id)
			var stu model.Student
			dao.DB.First(&stu,id)
			fmt.Printf("id : %d  name : %s  class : %s  score : %d\n",stu.Id , stu.Name , stu.Class , stu.Score)
		case 2:
			fmt.Println("使用姓名查询")
			var name string
			fmt.Println("请输入要查询学生的姓名：")
			fmt.Scanln(&name)
			var stu model.Student
			dao.DB.First(&stu,name)
			fmt.Printf("id : %d  name : %s  class : %s  score : %d\n",stu.Id , stu.Name , stu.Class , stu.Score)
		case 3:
			fmt.Println("使用班级查询")
			var class string
			fmt.Println("请输入要查询学生的班级：")
			fmt.Scanln(&class)
			order := "select id,`name`,class,score from students where class = ?"
			dao.Query(order,class)

		case 4:
			fmt.Println("使用成绩查询")
			var score string
			fmt.Println("请输入要查询的成绩：")
			fmt.Scanln(&score)
			order := "select id,`name`,class,score from students where score = ?"
			dao.Query(order,score)
		case 5:
			fmt.Println("退出查询")
			return
		default:
			fmt.Println("输入错误请重新输入：")
		}
	}
}

func reviseList(){
	fmt.Println("--------修改学生信息-------")
	fmt.Println("-------1.修改成绩---------")
	fmt.Println("-------2.修改班级---------")
	fmt.Println("-------3.退出修改---------")
}

func Revise(){
	var id ,number int
	fmt.Println("输入要修改学生的学号：")
	fmt.Scanln(&id)
	var stu model.Student
	dao.DB.First(&stu,id)
	fmt.Println("该学生的信息：")
	fmt.Printf("id : %d  name : %s  class : %s  score : %d\n",stu.Id , stu.Name , stu.Class , stu.Score)
	for{
		reviseList()
		fmt.Println("请输入你的选择：")
		fmt.Scanln(&number)
		switch number{
		case 1:
			var n int
			fmt.Println("请输入新的成绩：")
			fmt.Scanln(&n)
			s := "UPDATE students SET score = ? where id = ?"
			dao.Update(n , s , id)
			dao.AfterUpdate(id)
		case 2:
			var class string
			fmt.Println("请输入新的班级：")
			fmt.Scanln(&class)
			s := "UPDATE students SET class = ? where id = ?"
			dao.Update(class , s , id)
			dao.AfterUpdate(id)
		case 3:
			fmt.Println("退出修改")
			return
		}
	}
}


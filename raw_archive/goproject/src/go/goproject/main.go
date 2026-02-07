package main

import (
	"Aurora/struct/studentExercise/controller"
	"Aurora/struct/studentExercise/dao"
	"fmt"
	"os"
)

func main() {
	err := dao.InitMysql()
	if err != nil{
		//fmt.Println(err)
		return
	}
	defer dao.DB.Close()
	var number int
	//var c model.Class //用于管理学生
	////显示菜单
	//c.Stu = make(map[int]*model.Student , 20)
	for{
		controller.ShowMenu()
		fmt.Println("请输入你的选择：")
		fmt.Scanln(&number)
		switch number {
		case 1:
			fmt.Println("添加学生")
			controller.Add()
		case 2:
			fmt.Println("查询学生信息")
			controller.Find()
		case 3:
			fmt.Println("编辑学生信息")
			controller.Revise()
		case 4:
			fmt.Println("删除学生")
			controller.Delete()
		case 5:
			fmt.Println("展示学生列表")
			controller.ShowList()
		case 6:
			fmt.Println("退出系统")
			os.Exit(0)
		default :
			fmt.Println("输入错误请重新输入：")

		}

	}

}



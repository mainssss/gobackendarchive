package main

import (
	"fmt"
	"gocode/testroom/cilent/process"
	"os"
)

//定义两个变量，一个表示用户ID
var userId int

//一个表示用户密码
var userPwd string

var userName string

func main() {
	//接收用户选择
	var key int
	//判断是否进入下一个循环界面
	//var loop = true

	for true { //如果loop==true则继续运行
		fmt.Println("----------欢迎使用多人聊天系统----------")
		fmt.Println("\t\t1.登录聊天室")
		fmt.Println("\t\t2.注册用户")
		fmt.Println("\t\t3.退出系统")
		fmt.Println("请选择(1~3):")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登录聊天室")
			fmt.Println("请输入用户的ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)
			//完成登录
			//1.创建一个UserProcess实例
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
			//loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户昵称")
			fmt.Scanf("%s\n", &userName)
			//3.调用UserProcess实例，完成一个注册的请求
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
			//loop = false
		case 3:
			fmt.Println("退出系统")
			//loop = false
			os.Exit(0)
		default:
			fmt.Println("你的输入有误,请重新输入")
		}
	}

	//根据用户输入显示新的提示信息
	/*
		if key == 1 {
			//说明用户要登陆
			fmt.Println("请输入用户的ID:")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户的密码:")
			fmt.Scanf("%s\n", &userPwd)

			//因为使用了新的结构
			//先把登录的函数写到另外一个包内(login.go)
			login(userId, userPwd)

				if err != nil {
					fmt.Println("\n登录失败")
				} else {
					fmt.Println("\n登录成功")
				}

		} else if key == 2 {
			fmt.Println("进行用户注册...")
		}
	*/
}

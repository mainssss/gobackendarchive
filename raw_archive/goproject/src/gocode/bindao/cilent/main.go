package main

import (
	"fmt"
	"os"
)

//定义全局变量，一个id，一个密码
var userID int
var userPWD string

func main() {
	var key int
	var loop = true
	fmt.Println("---这是冰岛服务器...---")
	fmt.Scanln()
	fmt.Println("这是一款开发了一整个大学的项目,联合chatGPT开发...")
	fmt.Scanln()
	fmt.Println("准备好了就按下回车键吧")
	fmt.Scanln()
	for loop {
		fmt.Println("\t冰岛服务器")
		fmt.Println("\t1.登录")
		fmt.Println("\t2.注册")
		fmt.Println("\t3.退出")
		fmt.Println("请选择1~3:")
		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("登录聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			os.Exit(0)
		default:
			fmt.Println("你的输入有误,请重新输入")
		}
	}
	if key == 1 {
		//说明用户要登录
		fmt.Println("请输入用户ID")
		fmt.Scanf("%d\n", &userID)
		fmt.Println("请输入用户密码")
		fmt.Scanf("%s\n", &userPWD)
		err := Login(userID, userPWD)
		if err != nil {
			fmt.Println("登陆失败...")
		} else {
			fmt.Println("登陆成功")
		}
	} else if key == 2 {
		fmt.Println("注册用户")
	}
}

package main

import "fmt"

func main() {
	//声明变量，保存接收用户输入选项:
	key := ""
	//声明一个变量，控制是否退出for循环:
	loop := true
	//定义账户余额
	balanc := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	note := ""
	//定义一个变量，记录是否有收支行为
	flag := false
	//收支的详情使用字符串来记录
	//当有收支时，只需要对details进行拼接处理即可
	details := "收支\t账户余额\t收支金额\t说  明"
	//显示主菜单:
	for {
		fmt.Println("\n--------家庭收支软件--------")
		fmt.Println("        1 收支明细")
		fmt.Println("        2 登记收入")
		fmt.Println("        3 登记支出")
		fmt.Println("        4 退出软件")
		fmt.Println()
		fmt.Print("请选择(1~4):")

		fmt.Scanln(&key)

		switch key {
		case "1":
			fmt.Println("--------------当前收支明细--------------")
			if flag == true {
				fmt.Println(details)
			} else {
				fmt.Println("当前还没有任何收支状况...")
			}
		case "2":
			fmt.Println("本次收入金额:")
			fmt.Scanln(&money)
			balanc += money //修改账户余额
			fmt.Println("本次收入说明:")
			fmt.Scanln(&note)
			//将收入情况记入到details变量
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balanc, money, note)
			flag = true
		case "3":
			fmt.Println("登记支出")
			fmt.Println("本次支出金额:")
			fmt.Scanln(&money)
			//需要做一个必要的判断
			if money > balanc {
				fmt.Println("账户余额不足")
				break
			}
			balanc -= money
			fmt.Println("本次支出说明:")
			fmt.Scanln(&note)
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balanc, money, note)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗? y/n")
			fmt.Println("y:yes\nn:no")
			choice := ""
			for {
				fmt.Scanln(&choice)
				if choice == "y" || choice == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入 y/n")
			}
			if choice == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确选项")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出了软件")
}

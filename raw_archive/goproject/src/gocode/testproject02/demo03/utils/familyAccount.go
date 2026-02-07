package utils

import "fmt"

type FamilyAccount struct {
	//声明必要的字段.
	//声明变量，保存接收用户输入选项:
	key string
	//声明一个变量，控制是否退出for循环:
	loop bool
	//定义账户余额
	balanc float64
	//每次收支的金额
	money float64
	//每次收支的说明
	note string
	//定义一个变量，记录是否有收支行为
	flag bool
	//收支的详情使用字符串来记录
	//当有收支时，只需要对details进行拼接处理即可
	details string
}

//编写一个工厂模式的构造方法，返回一个*FamilyAccount的实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balanc:  10000.0,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户余额\t收支金额\t说  明",
	}
}

//将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("--------------当前收支明细--------------")
	if this.flag == true {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前还没有任何收支状况...")
	}
}

//将登记收入写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) income() {
	fmt.Println("本次收入金额:")
	fmt.Scanln(&this.money)
	this.balanc += this.money //修改账户余额
	fmt.Println("本次收入说明:")
	fmt.Scanln(&this.note)
	//将收入情况记入到details变量
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", this.balanc, this.money, this.note)
	this.flag = true
}

//将登记支出写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) pay() {
	fmt.Println("登记支出")
	fmt.Println("本次支出金额:")
	fmt.Scanln(&this.money)
	//需要做一个必要的判断
	if this.money > this.balanc {
		fmt.Println("账户余额不足")
		return
	}
	this.balanc -= this.money
	fmt.Println("本次支出说明:")
	fmt.Scanln(&this.note)
	this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", this.balanc, this.money, this.note)
	this.flag = true
}

//将退出系统写成一个方法
func (this *FamilyAccount) exit() {
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
		this.loop = false
	}

}

//给结构体绑定相应的方法:
//显示主菜单
func (this *FamilyAccount) MainMenu() {

	for {
		fmt.Println("\n--------家庭收支软件--------")
		fmt.Println("        1 收支明细")
		fmt.Println("        2 登记收入")
		fmt.Println("        3 登记支出")
		fmt.Println("        4 退出软件")
		fmt.Println()
		fmt.Print("请选择(1~4):")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确选项")
		}
		if !this.loop {
			break
		}

	}
}

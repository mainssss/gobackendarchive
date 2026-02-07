package process2

import (
	"encoding/json"
	"fmt"
	"gocode/testroom/common/message"
	"gocode/testroom/server/model"
	"gocode/testroom/server/utils"
	"net"
)

type UserProcess struct {
	//字段
	Conn net.Conn
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.先从mes中取出mes.Data，并且反序列化成registerMes:
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return
	}
	//2.先声明一个resMes,并完成赋值
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//在声明一个LoginResMes
	var registerResMes message.RegisterResMes
	//去redis数据库完成注册
	//1.使用model.MyUserDao 到redis去验证
	err = model.MyUserDao.Register(&registerMes.User)

	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误"
		}
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}
	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	//6.发送data 我们将其封装到writePkg
	//因为使用了分层模式(mvc),我们先创建一个Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码
	//1.先从mes中取出mes.Data，并且反序列化成LoginMes:
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("反序列化失败:", err)
		return
	}

	//2.先声明一个resMes,并完成赋值
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//在声明一个LoginResMes
	var loginResMes message.LoginResMes

	//我们需要到redis数据库去完成验证
	//1.使用model.MyUserDao 到redis去验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 300
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}

		//loginResMes.Code = 500
		//loginResMes.Error = "该用户不存在,请注册在使用..."
		//这里我们先测试成功，然后再根据具体错误信息
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登陆成功")
	}

	//如果用户id为1425919357，密码为123456，认为是正确，否则不合法
	/*
		if loginMes.UserId == 1425919357 && loginMes.UserPwd == "123456" {
			//合法
			loginResMes.Code = 200

		} else {
			//不合法
			loginResMes.Code = 500 //500状态码表示该用户不存在
			loginResMes.Error = "该用户不存在,请注册在使用..."
		}*/

	//3.将loginResMes序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes 进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("序列化失败:", err)
		return
	}

	//6.发送data 我们将其封装到writePkg
	//因为使用了分层模式(mvc),我们先创建一个Transfer 实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

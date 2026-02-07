package model

//定义一个用户的结构体

type User struct {
	//确定字段信息
	//为了能够序列化和反序列化成功，我们必须保证
	//用户信息的json字符串的key和结构体对应的tag 名字一致!!!
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

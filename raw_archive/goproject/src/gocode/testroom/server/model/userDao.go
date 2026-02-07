package model

import (
	"encoding/json"
	"fmt"

	"gocode/testroom/common/message"

	"github.com/garyburd/redigo/redis"
)

//我们在服务器启动后就初始化一个userDao的实例，
//把它做成全局的变量，在需要和redis操作时就直接使用即可
var (
	MyUserDao *UserDao
)

//定义一个UserDao结构体
//完成对User结构体的各种操作

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建一个userDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {

	userDao = &UserDao{
		pool: pool,
	}
	return
}

//在UserDao 应该提供哪些方法
//1.根据用户id返回一个user实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//通过给定的id 去redis查询用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误!
		if err == redis.ErrNil { //表示在users哈希中我们没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	//需要把res 反序列化成user实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err = ", err)
		return
	}
	return
}

//完成登录校验的函数
//1.Login
//2.如果用户id和密码都正确，则返回一个user实例
//3.如果用户id和密码有错，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//userDao连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//这时证明用户获取到.
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {

	//userDao连接池中取出一根连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err != nil {
		err = ERROR_USER_EXISTS
		return
	}

	//这时说明该Id在redis里面还没有，则可以完成注册
	data, err := json.Marshal(user) //序列化
	if err != nil {
		return
	}

	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户信息错误 err:", err)
		return
	}
	return
}

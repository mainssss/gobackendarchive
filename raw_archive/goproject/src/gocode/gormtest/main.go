package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GLOBAL_db *gorm.DB

func main() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:chenjie110@tcp(127.0.0.1:3306)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local",
		DefaultStringSize: 191,
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		// NamingStrategy: schema.NamingStrategy{
		// 	TablePrefix:   "gva_", // 表名前缀,'user'的表名应该是't_user'
		//	SingularTable: false,  // 使用单数表名,启用该选项,此时.'user'的表名应该是't_user',如果是false,则表是't_users'
		//},
		DisableForeignKeyConstraintWhenMigrating: true, //逻辑外键
	})

	// type user struct {
	// 	Name string
	// }
	// type userTwo struct {
	// 	Name string
	// }

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)           //连接池中最大的空闲连接数
	sqlDb.SetMaxOpenConns(100)          //连接池最大连接数量
	sqlDb.SetConnMaxLifetime(time.Hour) //连接池的最大可用时间
	GLOBAL_db = db
	//_ = db.AutoMigrate(&user{}) //自动创建表
	// m := db.Migrator()
	// //m.CreateTable(&user{})
	// if m.HasTable(&user{}) {
	// 	m.RenameTable(&user{}, &userTwo{}) //重新命名表,把前面的名字命名为后面

	// } else {

	// 	m.RenameTable(&userTwo{}, &user{})
	// }
	//fmt.Println(m.HasTable("gva_users"))  //是否存在表
	//fmt.Println(m.DropTable("gva_users")) //删除表
	TestUserCreate()

	//CreateTest()
	//FindTest2()
	TestUpdata()
	One2one()
}

func CreateTest() { //创建
	//GLOBAL_db.Omit()
	dbres := GLOBAL_db.Create(&[]TestUser{
		{Name: "陈杰", Age: 18},
		{Name: "陈1", Age: 19},
		{Name: "陈2", Age: 20},
		{Name: "陈3", Age: 21},
		{Name: "陈4", Age: 22},
	}) //.Select:只想创建的信息,omit则反之
	fmt.Println(dbres.Error, dbres.RowsAffected) //是否有报错和影响的条数
	if dbres.Error != nil {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")
	}
}

func FindTest() { //查询
	var user TestUser //被赋值结构体
	//GLOBAL_db.Model(&TestUser{}).First(&user) //按照第一条主键查询,不过得先创建一个结构体
	//dbres := GLOBAL_db.Model(&TestUser{}).First(&user, 14) //在后面加个int表示按主键14查询
	//.Last()拿最后一条数据
	//fmt.Println(errors.Is(dbres.Error, gorm.ErrRecordNotFound)) //把两种错误进行比较
	GLOBAL_db.Where("Name = ?", "陈杰").First(&user)      //使用string条件查询
	GLOBAL_db.Where(&TestUser{Name: "陈杰"}).First(&user) //使用结构体查询,与上述一样
	GLOBAL_db.Where(map[string]interface{}{             //使用map
		"Name": "陈杰",
	}).First(&user)
	fmt.Println(user)
}

type UserInfo struct {
	Name string
	Age  int
}

func FindTest2() {
	var user []TestUser
	var u []UserInfo //使用切片可以一次接收多条
	//GLOBAL_db.First(&user, "name = ?", "陈杰") //内联
	//GLOBAL_db.Find(&user)//找所有内容
	GLOBAL_db.Where("name <> ?", "陈杰").Find(&user)                    //排除陈杰查询其他
	GLOBAL_db.Where("name LIKE ?", "%陈%").Find(&user)                 //查询所有带陈的
	GLOBAL_db.Model(&TestUser{}).Where("name LIKE ?", "%陈%").Find(&u) //将&TestUser{}中的符合u的字段取出
	fmt.Println(u)
}

func TestUpdata() { //更新
	//updata 只更新你选择的字段
	//updatas 更新所有字段
	//save 无论如何都更新所有内容
	var user []TestUser
	GLOBAL_db.Model(&TestUser{}).Where("name LIKE ?", "%陈%").Update("name", "帅")
	// dbres := GLOBAL_db.Model(&TestUser{}).Where("name LIKE ?", "%帅%").Find(&users)
	// for k := range users {
	// 	users[k].Age = 18
	// }
	// dbres.Save(&users) //所有包含帅的Age都改为18
	GLOBAL_db.First(&user).Updates(TestUser{Name: "第一条", Age: 0})                  //以结构体进行更新时,零值不参与更新
	GLOBAL_db.Find(&user).Updates(map[string]interface{}{"Name": "第一条", "Age": 0}) //更改多条
}

func TestDelect() { //删除
	var user TestUser
	GLOBAL_db.Where("name = ?", "").Delete(&user)            //删除所有name为空的(软删除)
	GLOBAL_db.Unscoped().Where("name = ?", "").Delete(&user) //硬删除

}

type TestUser struct {
	gorm.Model
	Name string `gorm:"default:"qm"`
	Age  uint8  `gorm:"comment:年龄"`
}
type Model struct {
	UUID uint      `gorm:"primaryKey"` //gorm标签,主键
	Time time.Time `gorm:"column:my_time"`
}

func TestUserCreate() {
	GLOBAL_db.AutoMigrate(&TestUser{})
}

// belongsTo

type Dog struct {
	gorm.Model
	Name      string
	GirlGodID uint
	GirlGod   GirlGod
}

type GirlGod struct {
	gorm.Model
	Name string
}

func One2one() {

	g := GirlGod{
		Model: gorm.Model{
			ID: 1,
		},
		Name: "mm",
	}
	d := Dog{
		Model: gorm.Model{
			ID: 1,
		},
		Name:    "qiqi",
		GirlGod: g,
	}

	GLOBAL_db.AutoMigrate(&d)
}

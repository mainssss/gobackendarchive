package model
import "fmt"

type persoon struct{
	Name string
	age int  //首字母小写，其他包不能访问
}

//定义工程模式的函数，相当于构造器:
func NewPersoon(name string) *persoon{
	return &persoon{
		Name : name,
	}
}

//定义set和get方法，对age函数进行封装，因为在这个方法中可以加一系列的限制操作，确保被封装字段的安全合理性
func (p *persoon) SetAge(age int){
	if age >0 && age < 150{
		p.age = age
	}else{
		fmt.Println("年龄不在范围之内")
	}
}

func (p *persoon) GetAge()int{
	return p.age
}
package main
import "fmt"

type Student struct{
	Name string
	Age int
}

func main(){
	//方式1:按顺序赋值
	var s1 Student = Student{"小李",19}
	fmt.Println(s1)
	//方式2:按照指定类型赋值
	var s2 Student = Student{
		Name : "lili",
		Age : 20,
	}
	fmt.Println(s2)
	//方式3:想要返回结构体的指针类型
	var s3 *Student = &Student{"明明",26}
	fmt.Println(*s3)

	var s4 *Student = &Student{
		Name : "Nana",
		Age : 29,
	}
	fmt.Println(*s4)
}
package main
import "fmt"

type Student struct{
	Name string
}

//定义方法:
func (s Student) test01(){
	fmt.Println(s.Name)
}

//定义一个函数:
func method01(s Student){
	fmt.Println(s.Name)
}

func main(){
	//调用函数:
	var s Student = Student{"LiLi"}
	method01(s)

	//方法的调用:
	s.test01()
}
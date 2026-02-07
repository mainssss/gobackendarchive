package main
import "fmt"

type Student struct{
	Name string
}

//定义方法:
func (s Student) test01(){
	fmt.Println(s.Name)
}

func (s *Student) test02(){
	fmt.Println(s.Name)
}

//定义一个函数:
func method01(s Student){
	fmt.Println(s.Name)
}

func method02(s *Student){
	fmt.Println((*s).Name)
}

func main(){
	var s Student = Student{"LiLi"}
	s.test01()
	(&s).test01()//虽然用指针类型调用，但是还是按照值传递的形式
	s.test02()

}

// func main(){
// 	var s Student = Student{"LiLi"}
// 	method01(s)
// 	method02(&s)
// }
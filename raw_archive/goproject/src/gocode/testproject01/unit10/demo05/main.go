package main
import "fmt"
//定义Person结构体:
type Person struct{
	Name string
}

//给Person结构体绑定一个方法test:
func (p *Person) test(){ //参数名字随便起
	p.Name = "lulu"
	fmt.Println(p.Name)
}

func main(){
	//创建结构体对象:
	var p Person
	p.Name = "lili"
	fmt.Printf("p的地址是%p",&p)
	p.test()
	fmt.Println(p.Name)
}
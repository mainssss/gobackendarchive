package main
import "fmt"

//定义动物的结构体:
type Animal struct{
	Age int
	weight float32

}

//给Animal定义一个方法:喊叫:
func (an *Animal) Shout(){
	fmt.Println("大声喊叫")
}

//给Animal定义一个方法:自我展示:
func (an *Animal) showInfo(){
	fmt.Printf("动物的年龄：%v , 动物的体重: %v",an.Age,an.weight)
}

//定义结构体:Cat
type Cat struct{
	//为了复用性，体现继承的思维，嵌入匿名结构体:
	Animal
	Age int
}

func (c *Cat) showInfo(){
	fmt.Printf("~~~~动物的年龄：%v , 动物的体重: %v",c.Age,c.weight)
}

//对Cat绑定特有的方法:
func (c *Cat) scarth(){
	fmt.Println("猫")
}

func main(){
	//创建猫的结构体实例:
	// cat := &Cat{}
	// cat.Age = 3
	// cat.weight = 10.6
	// cat.Shout()
	// cat.showInfo()
	// cat.scarth()
	cat := &Cat{}
	cat.weight = 9.4
	cat.Age = 10
	cat.Animal.Age = 20
	cat.showInfo()
	cat.Animal.showInfo()
}
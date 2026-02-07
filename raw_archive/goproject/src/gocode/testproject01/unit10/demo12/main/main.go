package main
import (
	"fmt"
	"gocode/testproject01/unit10/demo12/model"
)
func main(){
	//创建person结构体的实例:
	p := model.NewPersoon("Lili")
	p.SetAge(200)

	fmt.Println(p.Name)
	fmt.Println(p.GetAge())
	fmt.Println(*p)
}
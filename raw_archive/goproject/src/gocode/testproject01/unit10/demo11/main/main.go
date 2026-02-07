package main
import (
	"fmt"
	"gocode/testproject01/unit10/demo11/model"
)

func main(){
	//跨包创建结构体Student的实例:
	//var s model.Student = model.Student{"丽丽",10}
	// s := model.student{"Lili",10}
	// fmt.Println(s)

	s := model.NewStudent("娜娜",20)
	fmt.Println(*s)
}
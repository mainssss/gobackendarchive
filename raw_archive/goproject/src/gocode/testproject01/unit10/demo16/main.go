package main
import "fmt"
type AIntface interface{
	a()
}
type BIntface interface{
	b()
}
type Stu struct{

}
func (s Stu)a(){
	fmt.Println("aaa")
}
func (s Stu)b(){
	fmt.Println("bbb")
}

func main(){
	var s Stu
	var a AIntface = s
	var b BIntface = s
	a.a()
	b.b()

}
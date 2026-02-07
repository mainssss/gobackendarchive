package main
import "fmt"
func main(){
	//定义一个map：
	b := make(map[int]string)
	//增加操作：
	b[20095452] = "张三"
	b[20095387] = "李四"
	b[20098833] = "王五"

	//获取长度:
	fmt.Println(len(b))
	
	//遍历：
	for k,v := range b{
		fmt.Printf("key为:%v,value为:%v\t",k,v)
	} 
	fmt.Println("--------------------------")

	//加深难度:
	a := make(map[string]map[int]string)
	//赋值操作：
	a["班级1"] = make(map[int]string,3)
	a["班级1"][20096677] = "张三"
	a["班级1"][20098833] = "丽丽"
	a["班级1"][20097722] = "菲菲"

	a["班级2"] = make(map[int]string,3)
	a["班级2"][20089911] = "小明"
	a["班级2"][20085533] = "小龙"
	a["班级2"][20087244] = "小飞"

	for k1,v1 := range a{
		fmt.Println(k1)
		for k2,v2 := range v1{
			fmt.Printf("学号为:%v,姓名为:%v\t",k2,v2)
		}
		fmt.Println()
	}
	

}
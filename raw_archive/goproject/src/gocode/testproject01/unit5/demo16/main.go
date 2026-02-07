package main
import (
	"time"
	"fmt"
)

func main(){
	//时间和日期的函数，需要导入time包，所以你获取当前时间，就要调用函数Now函数：
	now := time.Now()
	//Now()返回值是一个结构体，类型是：time.Time
	fmt.Printf("%v~~~~~ 对应的类型为%T \n",now,now)
	//2022-05-04 12:18:55.588249 +0800 CST m=+0.002998601~~~~~ 对应的类型为time.Time

	//调用结构体中的方法：
	fmt.Printf("年：%v \n",now.Year())
	fmt.Printf("月：%v \n",now.Month())//月：May
	fmt.Printf("月：%v \n",int (now.Month()))//月：5
	fmt.Printf("日：%v \n",now.Day())
	fmt.Printf("时：%v \n",now.Hour())
	fmt.Printf("分：%v \n",now.Minute())
	fmt.Printf("秒：%v \n",now.Second())

	fmt.Println("--------------------------------------")
	//Printf将字符串直接输出：
	fmt.Printf("当前年月日： %d-%d-%d   时分秒：  %d:%d:%d\n",now.Year(),now.Month(),
	now.Day(),now.Hour(),now.Minute(),now.Second())
	//Sprintf可以得到字符串，以便后续使用:
	datestr := fmt.Sprintf("当前年月日： %d-%d-%d   时分秒：  %d:%d:%d\n",now.Year(),now.Month(),
	now.Day(),now.Hour(),now.Minute(),now.Second())
	fmt.Println(datestr)
	//这个参数字符串各个数字必须是固定的，必须这样写
	datestr2 := now.Format("2006/01/02 15/04/05")
	fmt.Println(datestr2)
	//选择任意的组合都可以的，根据自己选择就可以(自己任意组合)。
	datestr3 := now.Format("2006 15:04")
	fmt.Println(datestr3)

}




package main
import "fmt"

func main(){
	//实现的功能:给出五个学生的成绩，求出总成绩和、平均数:
	//给出五个学生的成绩：----》数组存储
	//定义一个数组：
	var scores [5]int
	//将成绩存入数组:(循环+终端输入)
	for i := 0;i < len(scores);i++{//i:数组下标
		fmt.Printf("请录入第%d个学生的成绩",i + 1)
		fmt.Scanln(&scores[i])
	}
	//求和：
	//定义一个变量专门接收这个和：
	//sum := 0
	//for i := 0 ; i<len(scores) ; i++{
	//	sum += scores[i]
	//}
	//平均数:
	//avg := sum / len(scores)
	//将结果输出:
	//fmt.Printf("成绩的总和为：%v,成绩的平均数为：%v",sum,avg)

	//展示班级每个学生成绩:(数组遍历)
	//方式1：普通for循环:
	for i := 0;i < len(scores);i++{
		fmt.Printf("第%d个学生成绩为:%d\n",i+1,scores[i])
	}
	fmt.Println("-------------------------------------------")

	//方式2：for-range循环
	for key,value := range scores {
		fmt.Printf("第%d个学生的成绩为:%d\n",key + 1,value)
	}

	fmt.Println("-------------------------------------------")
	for _,value := range scores {
		fmt.Printf("学生的成绩为:%d\n",value)
	}
}
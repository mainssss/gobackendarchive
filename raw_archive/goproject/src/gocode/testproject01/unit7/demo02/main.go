package main
import "fmt"

func main(){
	//实现的功能:给出五个学生的成绩，求出总成绩和、平均数:
	//给出五个学生的成绩：----》数组存储
	//定义一个数组：
	var scores [5]int
	//将成绩存入数组:
	scores [0] = 95
	scores [1] = 91
	scores [2] = 39
	scores [3] = 60
	scores [4] = 21
	//求和：
	//定义一个变量专门接收这个和：
	sum := 0
	for i := 0 ; i<len(scores) ; i++{
		sum += scores[i]
	}
	//平均数:
	avg := sum / len(scores)
	//将结果输出:
	fmt.Printf("成绩的总和为：%v,成绩的平均数为：%v",sum,avg)
}
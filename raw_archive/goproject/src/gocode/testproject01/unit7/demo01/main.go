package main
import "fmt"

func main(){
	//实现的功能:给出五个学生的成绩，求出总成绩和、平均数:
	//给出五个学生的成绩
	score1 := 95
	score2 := 91
	score3 := 39
	score4 := 60
	score5 := 21
	//求和：
	sum := score1 + score2 + score3 + score4 +score5 
	//平均数:
	avg := sum / 5
	//将结果输出:
	fmt.Printf("成绩的总和为：%v,成绩的平均数为：%v",sum,avg)
}
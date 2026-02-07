package main
import "fmt"

func main(){
	//实现功能：根据给出的学生分数，判断学生的等级：
	// >=90 -----A
	// >=80 -----B
	// >=70 -----C
	// >=60 -----D
	// <60  -----E

	//方式1：利用if单分支实现：
	//定义一个学生的成绩：
	var score int = 18
	//对学生的成绩进行判定：
	// if score >= 90{
	// 	fmt.Println("您的成绩为A级别")
	// }
	// if score >= 80 && score < 90 {
	// 	fmt.Println("您的成绩为B级别")
	// }
	// if score >= 70 && score < 80 {
	// 	fmt.Println("您的成绩为C级别")
	// }
	// if score >= 60 && score < 70 {
	// 	fmt.Println("您的成绩为D级别")
	// }
	// if score <60 {
	// 	fmt.Println("您的成绩为E级别")
	// }

	//上面方式1利用多个单分支拼凑出多个选择，多个选择是并列的，依次从上而下顺序执行，即使走了第一个分支，那么其他分支也是需要判断的

	// //方式2：多分支：优点：如果已经走了一个分支，那么下面分支就不会再去判断执行了
	// if score >= 90{
	// 	fmt.Println("您的成绩为A级别")
	// } else if score >= 80 {//else隐藏：score < 90
	// 	fmt.Println("您的成绩为B级别")
	// } else if score >= 70 {//score < 80
	// 	fmt.Println("您的成绩为C级别")
	// } else if score >= 60 {//score < 70
	// 	fmt.Println("您的成绩为D级别")
	// } else {//score < 60
	// 	fmt.Println("您的成绩为E级别")
	// }//建议你保证else的存在，只有有了else才会真正 起到多选一 的效果

	if score > 10 {
		fmt.Println("aaa")
	} else if score > 6{
		fmt.Println("bbb")
	}
}

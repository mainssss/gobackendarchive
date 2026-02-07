package main
import "fmt"

type Student struct{
	Age int
}

type Stu Student

func main(){
	var s1 Student
	var s2 Stu
	s1 = Student(s2)
	fmt.Println(s1)
	fmt.Println(s2)
}
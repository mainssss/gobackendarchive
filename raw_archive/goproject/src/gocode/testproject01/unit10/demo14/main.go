package main
import "fmt"

type A struct{
	a int
	b string
}

type B struct{
	c int
	d string
	a int 
}

type C struct{
	A
	B
	int
}

type C1 struct{
	*A
	*B
	int
}

type D struct{
	a int
	b string
	c B //组合模式
}
func main(){
	d := D{10,"ooo",B{66,"ppp",99}}
	fmt.Println(d)
	fmt.Println(d.c.d)
	//构建C结构体实例:
	//c := C{A{10,"aaa"},B{20,"ccc",50},888}
	c := C{
		A{
			a : 10,
			b :"aaa"},
		B{
			c : 20,
			d : "ccc",
			a : 50},
		888}
	fmt.Println(c.b)
	fmt.Println(c.d)
	fmt.Println(c.A.a)
	fmt.Println(c.B.a)
	fmt.Println(c.int)

	c1 := C1{&A{10,"aaa"},&B{20,"ccc",50},888}
	fmt.Println(c1)
	fmt.Println(*c1.A)
	fmt.Println(*c1.B)
}
package main
import "fmt"

type fruits interface{
	pick()
	harvest()
}

type apple struct{
	name string
}

type orange struct{
	name string
}
func (*apple)pick(){
	fmt.Println("采摘苹果")
} 
func (*apple)harvest(){
	fmt.Println("收获苹果")
}
func (*orange)pick(){
	fmt.Println("采摘橘子")
} 
func (*orange)harvest(){
	fmt.Println("收获橘子")
}
func Newfruits(i int) fruits {
	switch i {
	case 1:
		return &apple{}
	case 2:
		return &orange{}
	}
	return nil
}


func main(){
	apple := Newfruits(1)
	apple.pick()
	apple.harvest()
	orange := Newfruits(2)
	orange.pick()
	orange.harvest()
}
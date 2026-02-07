package main

import (
	"fmt"
	"reflect"
)

//v的kind
func refV(b interface{}) {
	kind := reflect.ValueOf(b)
	fmt.Printf("kind  =%v\n", kind.Kind())
	intf := kind.Interface()
	fmt.Printf("interface=%T\n", intf)
	iV := intf.(float64)
	fmt.Printf("iV type=%v", iV)
}

//v的type
func rTyp(b interface{}) {
	rTyp := reflect.TypeOf(b)
	fmt.Printf("rTyp type =%v\n", rTyp)
}

func main() {
	var v float64 = 1.2
	refV(v)
	rTyp(v)
}

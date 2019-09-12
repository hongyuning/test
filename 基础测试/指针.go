package main

import (
	"fmt"
	"unsafe"
)
type student struct {

}
func (s *student)TEST(){
	fmt.Println("88888")
}

func main1 () {
	a := "894151512545qw1dwqfewfew"
	p := &a
	fmt.Printf("%T\n", p)
	fmt.Println(unsafe.Sizeof(p))
	fmt.Println(unsafe.Sizeof(a))
	aa:=student{}
	f:=aa.TEST
	f()
}
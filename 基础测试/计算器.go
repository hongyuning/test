package main

import "fmt"
type Father struct {
	n1,n2 int
}

type Calculate interface {
	GetResult() int
}

type Addition struct {
Father
}

func (add *Addition) GetResult() int {
	return add.n1 + add.n2
}
func Realize(calculate Calculate) int {
	return calculate.GetResult()
}

type SunTraction struct {
	Father
}

func (st *SunTraction) GetResult() int {
	return st.n1 - st.n2
}
type Call struct {

}
func (C *Call)CallRealize(option string,n1,n2 int){
	switch option {
	case "+":
		add := &Addition{Father{n1, n2}}
		fmt.Println(Realize(add))
	case "-":
		subtranction := &SunTraction{Father{
			n1: n1, n2: n2,}}
		fmt.Println(Realize(subtranction))
	}
}

func main3() {
	call:=Call{}
	call.CallRealize("-",9,8)
}

package main

import (
	"fmt"
)

func teat(a int) {
	fmt.Println("11111111")
}
func teat1(a int) {
	fmt.Println("22222222222")
}

type functest struct {
	start func(int2 int)
	stop  func(int2 int)
}

func (ft *functest) Callstart() {
	ft.stop(3)
}
func (ft *functest) Callstop() {
	ft.start(2)
}
func (ft *functest) Registerstart(f func(int)) {
	ft.start = f
}
func (ft *functest) Registerstop(f func(int)) {
	ft.stop = f
}

func main() {
	s := new(functest)
	s.Registerstart(teat)
	s.Registerstop(teat1)
	s.Callstart()
	s.Callstop()
}

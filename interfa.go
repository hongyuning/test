package main

type test interface {
	Registerstart(func(int))
	Registerstop(func(int))

	Callstart()
	Callstop()
}

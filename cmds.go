package main

import (
	"fmt"
	"os"
)

func  main (){
	cmds:=os.Args

	for k,v:=range cmds{
		fmt.Println(k,v)
	}

}

package main

import (
	"fmt"
	"math/big"
)

func main (){

	a:="1010"
	aaa:=[]byte("9")
	biga:=big.Int{}
	biga.SetString(a,2)
    bigaaa:=big.Int{}
    bigaaa.SetBytes(aaa)
   bb:=biga.Cmp(&bigaaa)
   fmt.Println(aaa)
fmt.Println(bb)
}

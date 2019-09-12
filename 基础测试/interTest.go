package main

import "fmt"

func te (in interface{}){
	fmt.Printf("%p   ,\n",in)
	fmt.Println(in)
	if in==nil{
		fmt.Println("11111")
	}else {
		fmt.Println("22222")
	}
}

func main (){
	var a *int
	te(a)

}
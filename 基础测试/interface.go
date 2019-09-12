package main

import (
	"fmt"
)

type Bird interface {
	Fly()
}

type maque struct {

}
func ( mq *maque)Fly(){
	fmt.Println("麻雀会飞....")
}
type yingwu struct {

}
func ( mq *yingwu)Fly(){
	fmt.Println("鹦鹉会飞....")
}
type Plane struct {

}
func ( mq *Plane)Fly(){
	fmt.Println("飞机会飞....")
}

func Flyfly(in Bird){
	in.Fly()
}
func main2 (){
	bird1:=&maque{}
	Flyfly(bird1)

}
package main

import (
	"bytes"
	"fmt"
	"unsafe"
)

func main1 (){
	buf := new(bytes.Buffer)
	var data = []interface{}{
		//uint16(61374),
		uint32(54),
		//uint8(254),
	}
	a:=make([]byte,len(data))
	for _, v := range data {
			//err := binary.Write(buf, binary.LittleEndian, v)
			//a:=v.([]byte)
			//n,err:=buf.Write(a)
			a=append(a,byte(v.(uint32)))
			//if err != nil {
			//	fmt.Println("binary.Write failed:", err,n)
			//}
		}
	n,err:=buf.Write(a)
	fmt.Println(n,err,len(a))
	fmt.Printf("%x", buf.Bytes())
}
func  main  (){
	var a uint
	var aa uint32
	var aaa uint64
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(aa))
	fmt.Println(unsafe.Sizeof(aaa))
}

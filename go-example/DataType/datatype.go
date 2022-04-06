package main

import "fmt"


func main(){
	var a int = 10
	var b uint = uint(a)
	var c float32 = float32(a)

	str:="ABC"
	bytes:=[]byte(str)
	str2:=string(bytes)
	
	println(bytes,str2)

	fmt.Println(a,b,c)
}
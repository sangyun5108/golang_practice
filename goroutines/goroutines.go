package main

import (
	"fmt"
	"time"
)

func main(){
	c:= make(chan string)
	people := [2]string{"a","b"}
	
	for _,person := range people {
		go show(person,c)
	}

	for i:=0; i<len(people); i++ {
		fmt.Println(<-c)
	}
}

func show(name string, c chan string){
	time.Sleep(time.Second*3)
	c <- "Hi" + name
}
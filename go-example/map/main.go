package main

import (
	"fmt"
	"scrapper/go-example/map/hello"
)

func main(){

	idMap := map[string]string{
		"A" : "a",
		"B" : "b",
		"C" : "c",
	}

	fmt.Println(idMap)

	map2:=make(map[string]string)

	map2["a"] = "A"

	_,exists := map2["a"]

	fmt.Println(exists)

	map3:=map[string]string{
		"1":"1",
		"2":"2",
		"3":"3",
	}

	for key,value := range map3 {
		fmt.Println(key,value)
	}

	fmt.Println(hello.Pop)
}
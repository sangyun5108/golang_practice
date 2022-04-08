package main

import "fmt"

func main(){
	var a = []int{1,2,3}
	
	s := a[0:3]

	s[0] = 10

	fmt.Println(s)
	fmt.Println(a)


	b:=make([]int, 5, 10)

	println(len(b), cap(b))

	g:=make([]int,0,3)

	for i:=1; i<=15; i++ {
		g = append(g,i)
		fmt.Println(len(g),cap(g))
	}

	fmt.Println(g)

	sliceA:=[]int{1,2,3}
	sliceB:=[]int{4,5,6}

	sliceA = append(sliceA, sliceB...)

	fmt.Println(sliceA)

	first := []int{0,1,2}
	second := make([]int,len(first), cap(first)*2)

	copy(second,first)

	fmt.Println(second)
	println(len(second),cap(second))
}
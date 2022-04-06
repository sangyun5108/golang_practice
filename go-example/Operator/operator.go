package main

import "fmt"

func main(){
	c := (3+5)*10/12
	fmt.Println(c)

	a := 10
	b := 20

	fmt.Println(a==b)
	fmt.Println(b<c)
	fmt.Println(a!=c)

	d := a==b || b>c && a<b

	fmt.Println(d)

	c = (a & b) << 1

	fmt.Println(c)

	a *= 10
	a += 29
	a >>= 2
	a |= 1

	var z int = 10
	var p = &z
	fmt.Println(p)
	fmt.Println(*p)
}
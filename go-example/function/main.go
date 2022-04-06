package main

import "fmt"
func main(){
	msg:="What are you doing?"

	say(&msg)

	answer("a","b","c")

	sum,total:=sum(1,2,3,4,5)

	fmt.Println(sum,total)
}

func say(msg *string){
	*msg = "I'm Codding"
}

func answer(msg ...string){
	for _,value:=range msg{
		println(value)
	}
}

func sum(nums ...int) (sum int,total int) {
	s:=0
	count:=0

	fmt.Println(nums)

	for _,n:=range nums {
		s+=n
		count++
	}

	return
}
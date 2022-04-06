package main

func main(){
	sum:=func(n...int) int {
		s:=0
		for _,value:= range n {
			s+=value
		}
		return s
	}

	result:=sum(1,2,3,4,5)
	println(result)

	answer:=calc(func(x int, y int) int {return x-y}, 10, 20)
	println(answer)
}

type calculator func(int,int) int

func calc(f calculator, a int, b int) int {
	result:= f(a,b)
	return result
}
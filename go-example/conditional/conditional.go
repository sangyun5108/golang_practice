package main

import "fmt"

func main(){
	k:=3

	if k == 1{
		println("one")
	}else if k == 2{
		println("two")
	}else {
		println("other")
	}

	max:=10
	i:=2

	if val:=i*2; val<max {
		println(val)
	}

	var name string
	var number = 1

	switch number {
	case 1:
		name = "What.."
	case 2:
		name = "is..."
	case 3:
		name = "your..."
	default:
		name = "I don' know"
	}

	println(name)

	grade(90)

	typeSwitch("Hello")

	check(1)

}

func grade(score int){
	switch {
	case score>=90:
		println("A")
	case score>=80:
		println("B")
	case score>=70:
		println("C")
	case score>=60:
		println("D")
	default:
		println("No Hope")
	}
}

func typeSwitch(i interface{}){
	switch i.(type){
	case int:
		println("int 형")
	case string:
		println("stirng 형")
	case bool:
		println("bool 형")
	default:
		fmt.Println("야무것도 아님")
	}

}

func check(val int ){
	switch val {
	case 1:
		println("1입니다")
		fallthrough
	case 2:
		println("2입니다")
		fallthrough
	case 3:
		println("3입니다")
		fallthrough
	case 4:
		println("4입니다")
		fallthrough
	
	default:
		println("마지막 도달")
	}
}
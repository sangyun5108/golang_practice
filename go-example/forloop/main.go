package main

func main(){
	n:=1

	for n<100 {
		n*=2

		println(n)
	}

	names:=[]string{"A","B","C"}

	for index,value:= range names{
		println(index,value)
	}

	a:=1

	for a<15 {
		if a==5 {
			a+= a
			continue 
		}

		a++

		if a>10{
			break
		}
	}

	if a == 11 {
		goto END
	}
	println(a)

	END:
		println("End")

	
	z:=0

	L1:
		for {
			if z == 0 {
				break L1
			}
		}

		print("L1 레이블 입니다!")
}
package main

func plusValue() func() int {
	i:=0

	return func() int {
		i++
		return i
	}
}

func main(){
	next := plusValue()

	println(next())
	println(next())
	println(next())

	anotherNext := plusValue()
	println(anotherNext())
	println(anotherNext())
}
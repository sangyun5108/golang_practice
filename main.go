package main

import (
	"fmt"

	"github.com/sangyunpark99/mydict"
)

func main() {
	// account1 := accounts.NewAccount("sangyun", 100)

	// account1.Deposit(100)
	// balance := account1.Balance()
	// fmt.Println("액수", balance)

	// // 에러처리
	// err := account1.Withdraw(5000)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	balance = account1.Balance()
	// 	fmt.Println("액수", balance)
	// }

	dict1 := mydict.Dictionary{"first": "First word"}

	value, err := dict1.Search("first")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}

	// add
	err2 := dict1.Add("first", "First word")
	if err2 == nil {
		fmt.Println(dict1)
	} else {
		fmt.Println(err2)
	}

	// update
	err3 := dict1.Update("asfadsf", "Just Word")
	if err3 == nil {
		fmt.Println(dict1)
	} else {
		fmt.Println(err3)
	}

	err4 := dict1.Delete("first")
	if err4 == nil {
		fmt.Println(dict1)
	} else {
		fmt.Println(err4)
	}

}

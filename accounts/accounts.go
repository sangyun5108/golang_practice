package accounts

import (
	"errors"
	"fmt"
)

var errorWithdraw = errors.New("Money is not enough")

type Account struct {
	owner   string
	balance int
}

// NewAccount creates  Account
func NewAccount(owner string, balance int) *Account {
	account := Account{owner: owner, balance: balance}

	return &account
}

// 입금 (method)
func (a *Account) Deposit(amount int) {
	fmt.Println("Deposit!")
	a.balance += amount
}

// 출금 (method)
func (a *Account) Withdraw(amount int) error {

	if a.balance < amount {
		return errorWithdraw
	}

	fmt.Println("WithDraw!")
	a.balance -= amount

	// null과 같은 역할
	return nil
}

// 돈 확인하기 (method)
func (a Account) Balance() int {
	return a.balance
}

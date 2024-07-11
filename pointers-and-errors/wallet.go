package main

type Wallet struct {
	balance int
}

func (w Wallet) Balance() int {
	return w.balance
}

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}

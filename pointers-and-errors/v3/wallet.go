package main

import (
	"errors"
)

var ErrSufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(mount int) {
	w.balance += mount
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Withdraw(mount int) error {
	if mount > w.balance {
		return ErrSufficientFunds
	}

	w.balance -= mount
	return nil
}

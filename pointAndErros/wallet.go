package main

import (
	"errors"
	"fmt"
)

// Learning note - using var in Go allows for a variable to become global in defined package
// in this case, main. So now everyone can see ErrInsufficentFunds. However, this is localized in
// this subfolder I think. Ran an experiment that proived I could not access ErrInsufficentFunds
// in the Hello.go file.
var ErrInsufficentFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {

	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficentFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

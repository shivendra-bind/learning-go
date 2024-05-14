package pointersnerrors

import (
	"errors"
	"fmt"
)

type Bitcoin float64

type Stringer interface {
	Strint() string
}

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = errors.New("insufficient balance")

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Ballance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

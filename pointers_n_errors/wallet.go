package pointersnerrors

import "fmt"

type Bitcoin float64

type Stringer interface {
	Strint() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Ballance() Bitcoin {
	return w.balance
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

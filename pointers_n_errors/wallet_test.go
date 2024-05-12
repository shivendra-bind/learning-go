package pointersnerrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		got := wallet.Ballance()
		if got != want {
			t.Errorf("got %s want %s ", got, want)
		}
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(15)
		got := wallet.balance

		if got != want {
			t.Errorf("got %s want %s ", got, want)
		}
	})
}

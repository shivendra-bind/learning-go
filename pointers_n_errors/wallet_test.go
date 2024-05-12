package pointersnerrors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	want := Bitcoin(10)
	got := wallet.Ballance()
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}

}

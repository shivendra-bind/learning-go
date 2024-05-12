package pointersnerrors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10))

	want := Bitcoin(10)
	got := wallet.Ballance()
	if got != want {
		t.Errorf("got %.2f want %.2f ", got, want)
	}

}

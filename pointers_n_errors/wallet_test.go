package pointersnerrors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	want := 10.0
	got := wallet.Ballance()
	if got != want {
		t.Errorf("got %.2f want %.2f ", got, want)
	}

}

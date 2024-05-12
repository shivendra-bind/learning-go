package pointersnerrors

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)

	want := 10
	got := wallet.Ballance()
	if got != want {
		t.Errorf("got %d want %d ", got, want)
	}

}

package pointersnerrors

import "testing"

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Ballance()
	if got != want {
		t.Errorf("got %s want %s ", got, want)
	}
}
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got error but didn't want one")
	}
}

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		got := wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(15)
		assertNoError(t, got)
		assertBalance(t, wallet, want)
	})

	t.Run("withdraw insufficient balance", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		want := Bitcoin(20)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, want)

	})
}

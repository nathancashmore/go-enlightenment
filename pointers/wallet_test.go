package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(BitCoin(10))

		assertBalance(wallet, t, BitCoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		startingBalance := BitCoin(30)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(BitCoin(100))

		assertBalance(wallet, t, BitCoin(30))
		assertError(err, t, ErrNoFunds)
	})
}

func assertError(got error, t *testing.T, want string) {
	t.Helper()
	if got == nil {
		t.Fatal(ErrNoFunds)
	}

	if got.Error() != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertBalance(wallet Wallet, t *testing.T, want BitCoin) {
	t.Helper()

	result := wallet.Balance()
	expected := want

	if result != expected {
		t.Errorf("got %s want %s", result, expected)
	}
}

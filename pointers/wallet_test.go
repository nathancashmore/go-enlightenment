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

	t.Run("Withdraw with funds", func(t *testing.T) {
		startingBalance := BitCoin(30)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(BitCoin(20))

		assertBalance(wallet, t, BitCoin(10))
		assertNoError(err, t)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := BitCoin(30)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(BitCoin(100))

		assertBalance(wallet, t, BitCoin(30))
		assertError(err, t, ErrNoFunds)
	})

	t.Run("Should display currency postfix", func(t *testing.T) {
		currency := BitCoin(10)

		got := currency.String()
		want := "10 BTC"

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func assertNoError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(got error, t *testing.T, want string) {
	t.Helper()
	if got == nil {
		t.Fatal(ErrNoFunds)
	}

	if got.Error() != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertBalance(wallet Wallet, t *testing.T, want BitCoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

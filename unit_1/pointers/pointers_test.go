package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	// Helper
	validate := func(wallet Wallet, want Bitcoin, t *testing.T) {
		t.Helper()
		got := wallet.balance
		if got != want {
			t.Errorf("wallet balance is %s but should be %s", got, want)
		}
	}

	// Helper
	validateError := func(err error, t *testing.T, want string) {
		t.Helper()
		if err == nil {
			t.Fatal("Wanted an error but got none")
		}

		if err.Error() != want {
			t.Errorf("Got %q want %q", err, want)
		}
	}

	t.Run("Deposit of bitcoin", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		var want Bitcoin = 10

		fmt.Printf("address of the pointers %p", &wallet.balance)

		validate(wallet, want, t)
	})

	t.Run("Withdrawal of bitcoin", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(20))

		err := wallet.Withdraw(10)

		var want Bitcoin = 10

		if err != nil {
			t.Errorf("Got error but didnt wanted")
		}
		validate(wallet, want, t)
	})

	t.Run("Insufficiant funds for withdrawl", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(30)
		validate(wallet, Bitcoin(20), t)

		validateError(err, t, ErrorMessage.Error())
	})
}

package main

import "testing"

// 我们想要一个钱包(wallet), 这个钱包有两个动作，一个是存钱(Deposit), 一个取钱(Withdraw),
// 还有一个是我们要知道现在钱包有多少钱了(Balance)
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, &wallet, 10)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 50}
		err := wallet.Withdraw(20)
		assertNoError(t, err)
		assertBalance(t, &wallet, 30)
	})

	t.Run("Withdraw out of money", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		err := wallet.Withdraw(100)
		assertError(t, err, ErrSufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet *Wallet, want int) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("want %v, but got %v", want, got)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatalf("did get any error but wanted ont")
	}
	if got != want {
		t.Fatalf("got %v want %v", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("don't want get err : %v", got)
	}
}


```go
ackage main

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(10)

	got := wallet.Balance()
	want := 10

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
```

### 编写最少得代码来运行测试并检查失败的测试输出

```go
package main

type Wallet struct {
}

func (w Wallet) Balance() int {
	return 0
}

func (w Wallet) Deposit(money int) int {
	return 0
}

```

### 编写足够的代码使得测试能通过

```go
package main

type Wallet struct {
	balance int
}

func (w Wallet) Balance() int {
	return w.balance
}

func (w Wallet) Deposit(amount int) {
	w.balance += amount
}
```

报错` wallet_test.go:14: got 0 want 10` 这很令人困惑

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

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}
```

### 这不太正确

报错` wallet_test.go:14: got 0 want 10` 这很令人困惑，我们的代码看起来是可以工作的
我们将新的金额添加到余额中去，然后余额方法应该返回当前状态

在Go中，请你调用函数或方法，参数将被复制
```
address of balance in Deposit is 0x1400000e160 
address of balance in test is 0x1400000e158
```

解释
```go
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

1.wallet := Wallet{}

这个代码创建了一个Wallet类型的原始对象，而不是指针。`wallet`是`Wallet`类型的值变量

2.wallet.Deposit(10)

这行代码调用了`Deposit`方法，取决于`Deposit`方法的定义，如果`Deposit`方法是这样定义的
```go
func (w Wallet)Deposit(amount int) {
	w.balance += amount
}
```
这里的`w Wallet`表示`Deposit` 方法接受的是`Wallet` 类型的变量。
当你调用`wallet.Deposit(10)`时，Go会创建`Wallet`的一个副本，并在这个副本上进行操作
因此不会改变原始对象的值

如果Deposit的方法定时是这样的
```go
func (w *Wallet)Deposit(amount int) {
	w.balance += amount
}
```

这里的`w *Wallet` 表示 Deposit 方法接受的是`Wallet` 类型的指针。
当你调用`wallet.Deposit(10)`时，Go会自动获取`wallet`的地址指针，
并通过这个地址修改原始对象的值

为了确保代码正确工作，我应该将`Deposit`方法定义定义为接受指针，这样他能修改原始对象
```go
type Wallet struct {
	balance int
}

func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}
```

这样，当你调用wallet.Deposit(10)时，原始对象`balance`会被正确修改，你的测试也会通过



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


### 重构

我们制作bit钱包，为了计算bit钱包写一个结构体有些矫枉过正，int就其工作方式很好，他是可描述的

Go允许你创建新的类型，语法是 type MyName OriginalType

```go
type Bitcoin int

type Wallet struct {
    balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
    w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}
```
```go
func TestWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))

	got := wallet.Balance()

	want := Bitcoin(10)

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
```
要创建Bitcoin你只需要语法Bitcoin(999)

这样做，我们创建了一个新的类型，并且可以在它们上面申明方法，当你在类型上想实现一些特定域的功能
会变得非常有用

```go
type Stringer interface {
	String() string
}
```

词接口在fmt包里面定义，可以让你打印与%s字符时使用类型
```go
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
```

就如你看到的，在类型上申明创建方法的语法与在结构上创建方法的语法相同

我们将测试新的字符串，他们将会被String()代替

```go
if got != want {
	t.Errorf("got %s want %s",got, want)
}
```

他的实际效果是
```
wallet_test.go:12: got 10 BTC want 11 BTC
```

这使得我们在测试的情况变得更加清楚

下一个要求是Withdraw函数

### 首先先写测试

```go
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
```

然后报错`got 20 BTC want 10 BTC`, 修复这个报错

```go
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}
```

### 重构

```go
func TestWallet(t *testing.T) {
	assetBalance := func(t *testing.T, wallet *Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assetBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assetBalance(t, &wallet, Bitcoin(20))
	})
}
```

如果你尝试Withdraw超过账户中的剩余余额，会发生什么情况，目前，我们的情况是假设没有透支情况

使用 Withdraw时如何发出问题信号

在Go语言，通常函数返回err提供调用者检查并才去行动

然我们在测试中尝试一下

### 先写测试

```go
    t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assetBalance(t, &wallet, startingBalance)

		if err == nil {
			t.Errorf("wanted an err but not didn't get one")
		}
	})
```

```go
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("oh no")
	}

	w.balance -= amount
	return nil
}
```

### 重构

我们为错误检查创建一个快速测试助手，以提高测试可读性

```go

	assetError := func(t testing.TB, err error) {
		t.Helper()
		if err == nil {
			t.Error("wanted an err but not didn't get one")
		}
	}

	t.Run("withdraw insufficient funds", func(t *testing.T) {
        startingBalance := Bitcoin(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Bitcoin(100))
    
        assetBalance(t, &wallet, startingBalance)
        assetError(t, err)
	})
```
# hello-world

传统上你的学习一门语言的第一个项目是hello-world
* 在你喜欢的任何地方创建文件夹
* 创建一个新的文件叫`hello.go` 并且将下面代码写入

```go
package main

import "fmt"

func main() {
	fmt.Print("hello world")
}
```
用`go run hello.go`尝试去运行它，真的神奇，这样就运行起来了

### 它是如何工作的
当你写Go语言的项目时，你会有main这个包并且定义main这个函数，`package`是将go语言组合在一起的方式

用`func`关键字定义了具有名称和正文的函数

用`import "fmt"` 我们了一个使用一个包含我们用的Println的函数的fmt包

### 如何做测试
你如何做测试，你最好将输出内容和输出函数分离开

因此，我们将这些问题分开，以遍容易测试

```go
package main

import "fmt"

func SayHello() string {
	return "hello, world"
}

func main() {
	fmt.Println(SayHello())
}

```
我们用`func` 创建了一个的函数，然后我们加入`string`这个关键词，这个表示这个函数返回string类型

然后创建一个新的文件叫hello_test.go用于我们测试我们的hello函数

```go
package main

import "testing"

func TestSayHello(t *testing.T) {
	got := SayHello()
	want := "hello, world"

	if got != want {
		t.Errorf("want get %s, but get %s", got, want)
	}
}

```

运行以下内容可以运行测试
```shell
go test
```

### 编写测试
编写测试就像写函数一样，只有几条规则
* 他需要文件名称为 xxx_test.go
* 测试函数的开头必须是Test
* 测试函数只有一个参数`t *testing.T`
* 为了使用`testing.T` 类型，你需要导入testing包，就像我们导入的fmt包一样

现在，你只需要知道你的t类型`*testing.T` 是你进入测试框架的钩子就可以了，这样你就可以做一些事情，
比如当你想失败的时候用t.Fail()


### go doc文档

Go另一个高质量的是文档，你可以通过`godoc -http=localhost:8000 `启动文档，

绝大多数标准库都有出色的文档和示例，可以通过  http://localhost:8000/pkg/testing/ 看到testing包里的内容

如果你没有godoc 命令，那么你可能用的是最新的Go，你可以用 `go install golang.org/x/tools/cmd/godoc@latest` 手动安装它

### hello, you
现在我们有了测试，我们可以安全的迭代我们的软件

在上一个示例中，我们在编写代码后编写了测试，以便你知道如何编写测试和申明函数。从现在起，我们将首先编写测试

下一个要求是我们要制定问候语的收件人

我们从测试中捕获这些要求开始，这是基本的测试驱动开发，使得我们能确保我们的测试实际上是在测试我们想要的东西，
回顾性的编写测试时，即使代码没有按照预期工作，测试也能通过

```go
package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("ning")
	want := "hello, ning"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

现在你运行`go test`,你应该有一个编译错误
```
./hello_test.go:6:18: too many arguments in call to Hello
    have (string)
    want ()
```


### 关于源代码管理的说明
此时，如果你使用的是源代码管理（你应该这么做）我会commit按照原样写代码，我们有测试支持的工作软件

不过，我不会推送到主线，因为我计划重构，在这一点上提交是很好的，以防止你在重构的时候遇到困境

### 常量
```go
const englishHelloPrefix = "Hello, "
```

我们可以重构我们的代码
```go
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	return englishHelloPrefix + name
}
```

重构后，并没有破坏任何内容

### 再次hello world

下一个要求是，我们的函数用空字符串调用时，它默认打印"Hello, World", 而不是"Hello, "

```go
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
```

```go
const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
```

重构不仅适用于生产代码，我们可以重构我们的测试代码区
```go
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```

我们已经将断言重构为一个函数，这减少了重复，提高了我们测试的可读性，我们需要传输`testing.T` 
以便我们可以在需要的时候告诉测试代码失败

使用`t.Helper()` 后，当测试失败时，错误信息将实际指向调用了`assertCorrectMessage` 的测试用了，
而不是`assertCorrectMessage` 函数本身，这大大提高了测试报告的可读性和调试效率


让我们再次回顾这个周期，
* 写测试
* 使得编译器通过
* 运行测试，查看他是否会失败，并检查错误消息是否有意义
* 编写足够的代码让测试通过
* 重构

表面是上看，这似乎很乏味，但坚持反馈循环很重要

如果不写测试，你就意味着你要通过运行软件手动检查代码，这将破坏你心流的状态，优势是长远看你并不会节省自己的时间

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
我们

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
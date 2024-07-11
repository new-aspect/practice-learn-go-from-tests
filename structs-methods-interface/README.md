### 重构

我们的代码已经完成了这项工作，但没有任何矩形的明确内容。粗心的开发人员可能会尝试向
这些函数提供三角形的宽度和高度，但没有意识到他们会返回错误答案。

我们可以给函数制定具体的名称，例如`RectangleArea`。 一个更简洁的方案是定义我们自己的类型，
成为`Rectangle`。 它为我们封装了这个概念

```go
type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Width * rectangle.Height
}
```

我希望你同意将Rectangle传递函数可以更清楚的反馈我们的意图，但是使用结构的好处有很多，我们后续
继续说


### 重构
```go
func TestArea(t *testing.T) {

	check := func(t *testing.T, got, want float64) {
		t.Helper()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		want := 72.0
		check(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793
		check(t, got, want)
	})
}
```

我们的测试有些重复，

我们要做的是获取形状的合集，对他们调用 Area() 方法，然后检查结果

我们希望能够编写出可以将Rectangle和Circle传递给checkArea的函数

我们可以使用接口实现这个意图

在Go的静态类型语言中，接口是一个强大的概念，因为他们允许你创建可以与不同类型一起使用的函数，
并创建高度解耦的代码，同时仍能保持类型安全

```go
type Shape interface {
    Area() float64
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})
}
```

### 解耦
请注意，我们的助手不关心形状是Rectangle是Circle还是Triangle。通过声明接口，帮助器和具体类型
分离，并且完成其工作所需要的方法

这种接口声明在软件设计中非常重要，我们稍后在后面部分更详细介绍

### 进一步重构

现在你对结构有了一些了解，我们可以介绍"表驱动测试"

当你想用相同方式测试测试用例表的时候，表测试相当的有用

```go
func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 7.0}, 84.0},
		{Circle{6.0}, 113.09733552923255},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
```

这里面有一个新的语法是创建一个匿名结构，areaTest。我们是通过带有两个字段的[]struct来声明一个
结构体切片，即shape和want。然后我们用我们的案例填充切片

然后哦我们对其他任何切片一样迭代它们，使用结构字符来运行我们的测试

表驱动测试也许是你工具箱里面一个很棒的项目，但请确保你需要测试中的额外的噪音，

### 重构
实现非常好，但我们仍然需要一些改进

当你扫描这个
```
{Rectangle{12.0, 7.0}, 84.0},
{Circle{6.0}, 113.09733552923255},
{Triangle{12.0, 7.0}, 42},
```
的时候不知道这串数字是什么意思，你的目标应该让你的测试易于理解。

到目前为止，你只看到了结构体MyStruct{val1, val2}的用法，你可以选择命名
```
{shape: Rectangle{Width: 12.0, Height: 7.0}, want: 84.0},
{shape: Circle{Radius: 6.0}, want: 113.09733552923255},
{shape: Triangle{Width: 12.0, Height: 7.0}, want: 42},
```

### 确保你的测试输出有帮助
还记得我们之前实现`Triangle` 的测试失败吗？它打印了 struct_test.go:20: got 48 want 42

我们知道这是与`Triangle`相关，因为只有我们在使用它，但是如果是20种情况之一遛进应该怎么办，
开发人员如何知道是哪个案例失败了，这对开发人员来说不是一个好的体验，他们必须手动查看案例才知道
是哪个案例失败了

我们可以将错误信息改为%#v got %g want %g 。%#v格式字符将打印出我们的结构以及字符的值，
因此开发人员可以一目了然的看到正在测试的属性。

为了进一步提高测试的可读性，我们可以将want字段重命名为更具描述性的内容，例如hasArea

表测试驱动的另一个技巧使用t.run并命名测试用例。

通过将每个案例包装在t.Run中，你将会在失败时获得更清晰的测试输出，应为它将答应名称
```
TestArea/Triangle (0.00s)
        struct_test.go:22: got 48 want 42
```

你也可是使用go test -run TestArea/Rectangle表运行特定测试

这是很重要的一章，因为我们现在开始定义我们自己的类型。在像Go语言这样的静态型语言中，能够设置
构建易于理解、组合和测试的软件至关重要。

接口是一个好东西，可以将复杂性隐藏到系统的其他部分之外，在我们的程序中，我们的测试帮助程序代码
不需要知道它断言的确切形状，只需要知道如何询问它的确切面积

随着你对Go的越来越熟悉，你将看到接口和标准库的真正优势。您将了解标准库中定义的接口，这些接口随处
使用，并且通过针对您自己的类型实现它们，宁可以非常快速的实现许多出色的功能


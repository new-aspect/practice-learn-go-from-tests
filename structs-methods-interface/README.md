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
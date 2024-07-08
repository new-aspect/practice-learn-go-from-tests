
执行这个可以看到测试覆盖率
```shell
 go test -cover
```

输出
```
PASS
coverage: 100.0% of statements
ok      practice-learn-go-from-tests/arrays-and-slices  0.660s

```

got != want 会编译报错 invalid operation: got != want (slice can only be compared to nil)
因为Go语言不允许你用相等运算符，你可以写一个函数迭代每个got和want切片并检查他们的值，但是为了方便起见，
我们可以使用reflect.Equal

```shell
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
  
	// 在Go 1.21 以后，支持slice.Equal 函数查看
	if !slices.Equal[[]int](got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
```

注意DeepEqual是类型不安全的，也就是你用,这样对比也能编译通过，但是没有意义
```go
got := SumAll([]int{1, 2}, []int{0, 9})
want := "bob"

if !reflect.DeepEqual(got, want) {
    t.Errorf("got %v want %v", got, want)
}
```

### 切片可以切
切片可以切切片，语法是slice[low:high]。如果省略: 一侧的值，它将捕获改侧的的所有内容，在我们
的例子中，我们用numbers[1:]说"从1到末尾"，你可能希望花一些时间围绕切片编写其他测试，并尝试
使用切片运算符更熟悉的运用它

```go
func ExampleSliceLow() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(numbers[1:])
	// Output: [2 3 4 5 6 7 8 9]
}

func ExampleSliceHigh() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(numbers[:1])
	// Output: [1]
}
```

### 运行报错
```
panic: runtime error: slice bounds out of range [recovered]
    panic: runtime error: slice bounds out of range
```
不好了，虽然测试已经编译，但是它存在运行时错误，

编译时运行错误是我们的朋友，因为它能帮助我们编写可以运行的软件

运行时错误是我们的敌人，因为它会影响我们的用户

编写足够的代码使其通过
```go
func SumAllTails(numbersToTails ...[]int) []int {
	var sumTails []int

	for _, numbers := range numbersToTails {
		if len(numbers) == 0 {
			sumTails = append(sumTails, 0)
			continue
		}
		tail := numbers[1:]
		sumTails = append(sumTails, Sum(tail))
	}
	return sumTails
}
```


### 重构
我们测试断言的时候有一些重复代码，所以我们将他们提到一个函数里面
```go

func TestSumAllTails(t *testing.T) {

    checkSum := func(t *testing.T, got, want []int) {
        t.Helper()
        if !slices.Equal[[]int](got, want) {
            t.Errorf("want %v got %v", want, got)
        }
    }
    
    t.Run("make the sums of some slices", func(t *testing.T) {
        got := SumAllTails([]int{1, 2}, []int{0, 9})
        want := []int{2, 9}
        checkSum(t, got, want)
    })
    
    t.Run("safely sum empty slices", func(t *testing.T) {
        got := SumAllTails([]int{}, []int{3, 4, 9, 5})
        want := []int{0, 18}
        checkSum(t, got, want)
    })
}
```

我们像平常一样创建一个新函数checkSums，在本例中，我们展示了一种新技术，将函数变量分配给变量，
它看起来很奇怪，但是这样讲变量分配给string和int没有什么不同，函数实际上也是值
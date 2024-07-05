
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
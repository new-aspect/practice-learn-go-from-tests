
### Benchmarking 基准测试
```go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
```

你会看到，这段代码与测试代码非常类似。

`testing.B` 让你可以访问一个 `b.N` 的变量

当基准测试代码执行的时候，他会运行 `b.N` 并测试所花费的时间

运行基准测试，你可以用`go test -bench.` 或者再Windows Powershell中，使用`go test -bench="."`

输出示例
```shell
goos: darwin
goarch: arm64
pkg: practice-learn-go-from-tests/iteration
BenchmarkRepeat-8       14705896                82.03 ns/op
PASS
ok      practice-learn-go-from-tests/iteration  4.244s
```

这里的 82.03 ns/op 意思是运行的平均时间是 136 纳秒，为了运行这个测试，运行了 14705896 次


### 用StringBuild提升速度
使用strings.Builder可以高效的构建字符串，避免反复的拷贝
```go
func Repeat(a string, iterate int) string {
	var builder strings.Builder
	builder.Grow(len(a) * iterate) // 预先分配足够的内存
	for i := 0; i < iterate; i++ {
		builder.WriteString(a)
	}
	return builder.String()
}
```

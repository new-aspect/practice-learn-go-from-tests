在数组和切片中，你已经了解如何按照顺序存储值，现在，我们研究一种通过key存储项目并快速的找到的方法

map允许你类似于字典的方式存储项目，您可以将key视为单词，将value视为定义，还有什么比构建自己的字典
更好的方式来了解map呢？

### 先写测试

```go
func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")
	want := "this is just a test"
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}
```

```go
func Search(dictionary map[string]string, search string) string {
	return dictionary[search]
}
```

###重构
```go
func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just a test"}
	got := Search(dictionary, "test")
	want := "this is just a test"
	assertString(t, got, want)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

```

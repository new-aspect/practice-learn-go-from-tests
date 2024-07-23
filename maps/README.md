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

### 使用自定义类型
我们可以围绕map创建一个新的类型，并用Search成为一个新的字典方法

### 先写测试
我们有一个很好的办法搜索字典，但是，我们无法向字典添加新单词

```go
func TestAdd(t *testing.T) {
    directory := Directory{}
    directory.Add("test", "this is a simple test")
    got := directory.Search("test")
    assertString(t, got, "this is a simple test")
}
```

实现Add

```go
func (d Directory) Add(work, content string) {
	d[work] = content
}
```

### 指针, 副本
map的有趣属性是你可以修改他们而不需要将其地址传递（例如&myMap）

地图的一个问题是它可以是nil值，nil值在阅读时类似空映射，但尝试写入nil值会造成恐慌
因此，你永远都不应该初始化nil变量为map
```go
var m map[string]string
```
相反，你可以初始化一个空地图或使用make关键创造map
```go
var directory = map[string]string{}
// OR
var directory = make(map[string]string)
```

### 重构
将这个里面的字符串以变量的名称命名
```go
func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is a test")
	got := dictionary["test"]
	want := "this is a test"
	if want != got {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is a test")
	got := dictionary.Search("test")
	want := "this is a test"
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
```

将检查部分逻辑
```go
func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	word := "test"
	definition := "this is a test"

	dictionary.Add("test", "this is a test")

	assertDefinition(t, dictionary, word, definition)
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find add word: ", err)
	}
	assertString(t, got, definition)
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
```
我们为单词定义创建了变量，并将定义断言移转至自己的辅助函数

我们的Add函数看起来不错，但是我们没有考虑添加已经存在的值会发生什么

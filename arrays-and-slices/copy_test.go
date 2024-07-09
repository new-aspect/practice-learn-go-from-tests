package main

import "fmt"

// 说明x和y公用相同的底层数组，而z是一个独立的副本
func ExampleCopy() {
	x := [3]string{"张三", "李四", "王五"}

	y := x[:] // slice "y" points to the underlying array "x"

	z := make([]string, len(x))
	copy(z, x[:])

	y[1] = "Lisi"

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)

	// output: [3]string [张三 Lisi 王五]
	//[]string [张三 Lisi 王五]
	//[]string [张三 李四 王五]
}

// ExampleCopyToFreeMemory
// 1. 切片是对底层数组的引用，因此对贴片的修改会影响到原数组
// 2. 使用`copy`函数可以创建贴片的副本，从而使修改副本不会影响到原始数据。
// 3. 在处理大切片时，复制切片的部分数据到新切片，可以帮助释放不再需要的内存，从而允许垃圾回收器回收这部分内存。
func ExampleCopyToFreeMemory() {
	a := make([]int, 1e6)
	b := a[:2] // b:=a[:2]创建一个切片`b`, 他的长度为2，但共享`a`的底层数组

	c := make([]int, len(b)) //通过创建贴片`c`并拷贝`b`的内容，可以保持原始贴片`a`可以被垃圾回收（前提是没有其他应用`a`的地方）
	copy(c, b)
	fmt.Println(c)
	// output:[0 0]
}

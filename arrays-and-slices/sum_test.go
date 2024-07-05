package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15
		if want != got {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6
		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}
	// got != want 会编译报错 invalid operation: got != want (slice can only be compared to nil)
	// 因为Go语言不允许你用相等运算符，你可以写一个函数迭代每个got和want切片并检查他们的值，但是为了方便起见，
	// 我们可以使用reflect.Equal
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
	// 注意DeepEqual是类型不安全的，也就是你用
	// got := SumAll([]int{1, 2}, []int{0, 9})
	//	want := "bob"
	//
	//	if !reflect.DeepEqual(got, want) {
	//		t.Errorf("got %v want %v", got, want)
	//	}
	// 这样对比也能编译通过，但是没有意义

	// 在Go 1.21 以后，支持slice.Equal 函数查看
	if !slices.Equal[[]int](got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

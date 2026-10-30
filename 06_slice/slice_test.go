package demo

import (
	"slices"
	"testing"
)

// 数组长度固定，而且是值类型：赋值会复制整个数组
func TestArrayIsValue(t *testing.T) {
	a := [3]int{1, 2, 3}
	b := a // 复制了一份
	b[0] = 100

	if a[0] != 1 {
		t.Error("修改 b 不应该影响 a")
	}
}

// 切片 = 指向底层数组的指针 + len + cap
func TestMake(t *testing.T) {
	s := make([]int, 2, 5) // len=2, cap=5

	t.Log("s =", s, "len =", len(s), "cap =", cap(s))

	if len(s) != 2 || cap(s) != 5 {
		t.Error("len/cap 不对")
	}
}

// append 会自动扩容；一定要把返回值接回来
func TestAppend(t *testing.T) {
	s := []int{1, 2}
	s = append(s, 3)

	t.Log("s =", s, "len =", len(s), "cap =", cap(s))

	if len(s) != 3 {
		t.Errorf("期望长度 3，实际 %d", len(s))
	}
}

// 坑 1：子切片和原切片共享底层数组，改一个另一个也会变
func TestShareArray(t *testing.T) {
	a := []int{1, 2, 3}
	b := a[0:2]
	b[0] = 100

	t.Log("a =", a, "b =", b)

	if a[0] != 100 {
		t.Errorf("期望 a[0] 被改成 100，实际 %d", a[0])
	}
}

// 坑 2：对子切片 append，cap 还够的话会直接覆盖原切片的元素
func TestAppendOverwrite(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := a[:2]        // len=2, cap=4
	b = append(b, 99) // 没有扩容，直接写进了 a[2]

	t.Log("a =", a, "b =", b)

	if a[2] != 99 {
		t.Error("a[2] 应该被覆盖成 99")
	}

	// 解决办法：用三下标切片 a[low:high:max] 限制 cap，append 时就会强制复制
	c := []int{1, 2, 3, 4}
	d := c[:2:2]
	d = append(d, 99)

	if c[2] != 3 {
		t.Error("c 不应该被修改")
	}
	_ = d
}

// 想要一份独立的副本，就用 copy 或 slices.Clone
func TestCopy(t *testing.T) {
	a := []int{1, 2, 3}
	b := make([]int, len(a))
	copy(b, a)
	c := slices.Clone(a)
	b[0], c[0] = 100, 200

	if a[0] != 1 {
		t.Error("a 不应该被修改")
	}
}

// nil 切片和空切片：len 都是 0，都可以直接 append
func TestNilSlice(t *testing.T) {
	var s []int  // nil
	e := []int{} // 空切片，但不是 nil

	s = append(s, 1)

	if e == nil || len(s) != 1 {
		t.Error("nil / 空切片行为不对")
	}
}

// slices 标准库包（Go 1.21+）提供了很多常用操作
func TestSlicesPackage(t *testing.T) {
	s := []int{3, 1, 2}
	slices.Sort(s)

	if !slices.Equal(s, []int{1, 2, 3}) || !slices.Contains(s, 2) || slices.Index(s, 3) != 2 {
		t.Error("slices 包结果不对", s)
	}
}

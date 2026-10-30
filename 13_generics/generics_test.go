package demo

import (
	"cmp"
	"slices"
	"strconv"
	"testing"
)

// 泛型函数：[T any] 是类型参数，any 表示任何类型都可以
func Map[T, U any](s []T, f func(T) U) []U {
	out := make([]U, 0, len(s))
	for _, v := range s {
		out = append(out, f(v))
	}
	return out
}

func TestGenericFunc(t *testing.T) {
	got := Map([]int{1, 2, 3}, strconv.Itoa) // T=int, U=string 是编译器自动推断出来的

	if !slices.Equal(got, []string{"1", "2", "3"}) {
		t.Error("Map 结果不对", got)
	}
}

// 类型约束：限制 T 只能是哪些类型。~int 表示"底层类型是 int"的类型也可以
type Number interface {
	~int | ~int64 | ~float64
}

func Sum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

type MyInt int

func TestConstraint(t *testing.T) {
	if Sum([]int{1, 2, 3}) != 6 || Sum([]float64{1.5, 2.5}) != 4 || Sum([]MyInt{1, 2}) != 3 {
		t.Error("Sum 结果不对")
	}
}

// cmp.Ordered：标准库里定义好的"可以比较大小"的约束
func MaxOf[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func TestOrdered(t *testing.T) {
	if MaxOf(1, 2) != 2 || MaxOf("a", "b") != "b" {
		t.Error("MaxOf 结果不对")
	}
	// 其实 Go 1.21+ 已经内置了 max / min 函数
	if max(1, 5, 3) != 5 {
		t.Error("内置 max 不对")
	}
}

// 泛型类型：一个什么类型都能装的栈
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) { s.items = append(s.items, v) }

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if len(s.items) == 0 {
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

func TestGenericType(t *testing.T) {
	var s Stack[string]
	s.Push("a")
	s.Push("b")

	v, _ := s.Pop()
	if v != "b" {
		t.Errorf("期望 b，实际 %s", v)
	}
}

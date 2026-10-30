package demo

import (
	"slices"
	"testing"
)

// defer 会在函数返回前执行；多个 defer 按"后进先出"的顺序执行
func deferOrder() []int {
	var out []int
	func() {
		for i := range 3 {
			defer func() { out = append(out, i) }()
		}
	}()
	return out
}

func TestDeferOrder(t *testing.T) {
	got := deferOrder()

	if !slices.Equal(got, []int{2, 1, 0}) {
		t.Errorf("期望 [2 1 0]，实际 %v", got)
	}
	// 最常见的用法：f, _ := os.Open(...); defer f.Close()
}

// 坑：defer 后面函数的"参数"在写 defer 的那一刻就求值了，而不是在执行时
func TestDeferArgsEvaluated(t *testing.T) {
	var got int
	func() {
		x := 1
		defer func(v int) { got = v }(x) // 这里 x=1 已经被传进去了
		x = 2
		_ = x
	}()

	if got != 1 {
		t.Errorf("期望 1，实际 %d", got)
	}
}

// defer 可以修改命名返回值（常用于统一处理错误）
func double() (result int) {
	defer func() { result *= 2 }()
	return 5 // 先把 result 设成 5，然后执行 defer，result 变成 10
}

func TestDeferNamedReturn(t *testing.T) {
	if double() != 10 {
		t.Errorf("期望 10，实际 %d", double())
	}
}

// panic 用于"程序没法继续"的严重错误；recover 只能在 defer 里使用，用来接住 panic
func safeDivide(a, b int) (result int, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			ok = false
		}
	}()
	return a / b, true // b == 0 时会 panic: integer divide by zero
}

func TestRecover(t *testing.T) {
	if _, ok := safeDivide(1, 0); ok {
		t.Error("除以 0 应该被 recover 接住，返回 ok=false")
	}
	if r, ok := safeDivide(6, 3); !ok || r != 2 {
		t.Error("正常除法不对")
	}
	// 经验：普通错误用 error 返回值，不要拿 panic 当异常来用
}

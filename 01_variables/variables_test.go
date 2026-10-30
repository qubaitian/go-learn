package demo

import "testing"

// 三种声明变量的方式
func TestDeclare(t *testing.T) {
	var a int = 1 // 完整写法
	var b = 2     // 类型推断
	c := 3        // 短声明，只能在函数内用

	t.Log("a, b, c =", a, b, c)

	if a+b+c != 6 {
		t.Errorf("期望 6，实际得到 %d", a+b+c)
	}
}

// 没赋值的变量会有"零值"，Go 里不存在"未初始化"的变量
func TestZeroValue(t *testing.T) {
	var i int
	var s string
	var ok bool
	var p *int

	t.Logf("int=%d string=%q bool=%v pointer=%v", i, s, ok, p)

	if i != 0 || s != "" || ok || p != nil {
		t.Error("零值不符合预期")
	}
}

// 多重赋值：交换两个变量不需要临时变量
func TestSwap(t *testing.T) {
	x, y := 1, 2
	x, y = y, x

	if x != 2 || y != 1 {
		t.Errorf("交换失败 x=%d y=%d", x, y)
	}
}

// 常量 + iota：iota 在每个 const 块里从 0 开始，每行 +1
func TestConstIota(t *testing.T) {
	const Pi = 3.14

	type Weekday int
	const (
		Sunday  Weekday = iota // 0
		Monday                 // 1，自动沿用上一行的表达式
		Tuesday                // 2
	)

	t.Log("Pi =", Pi, "Sunday/Monday/Tuesday =", Sunday, Monday, Tuesday)

	if Tuesday != 2 {
		t.Errorf("期望 Tuesday=2，实际 %d", Tuesday)
	}
}

// 坑：内层作用域里用 := 会创建一个新变量，把外层的"遮住"
func TestShadowing(t *testing.T) {
	x := 1
	if true {
		x := 2 // 这是一个新的 x，不是修改外层的 x
		t.Log("内层 x =", x)
	}

	if x != 1 {
		t.Errorf("外层 x 应该还是 1，实际 %d", x)
	}
}

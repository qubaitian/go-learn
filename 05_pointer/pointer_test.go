package demo

import "testing"

// & 取地址，* 取值
func TestPointerBasic(t *testing.T) {
	x := 10
	p := &x // p 的类型是 *int
	*p = 20 // 通过指针修改 x

	if x != 20 {
		t.Errorf("期望 20，实际 %d", x)
	}
}

// Go 的函数参数全部是"值传递"（复制一份）；想修改调用方的变量，要传指针
func setValue(n int)    { n = 100 }
func setPointer(n *int) { *n = 100 }

func TestPassByValue(t *testing.T) {
	a, b := 1, 1
	setValue(a)
	setPointer(&b)

	t.Log("a =", a, "b =", b)

	if a != 1 || b != 100 {
		t.Error("值传递 / 指针传递结果不对")
	}
}

// new(T) 分配一个零值的 T，返回 *T
func TestNew(t *testing.T) {
	p := new(int)
	*p++

	if *p != 1 {
		t.Errorf("期望 1，实际 %d", *p)
	}
}

// 返回局部变量的指针是安全的（编译器会把它放到堆上），不像 C 那样会出错
func newCounter() *int {
	c := 0
	return &c
}

func TestReturnLocalPointer(t *testing.T) {
	p := newCounter()
	*p = 5

	if *p != 5 {
		t.Error("应该是 5")
	}
	// 注意：Go 没有指针运算，p++ 这种写法编译不过
	// 坑：对 nil 指针取值 (*p) 会 panic
}

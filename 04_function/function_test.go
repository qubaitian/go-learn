package demo

import "testing"

// 多返回值：Go 最常见的写法是 (结果, error)
func divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

func TestMultiReturn(t *testing.T) {
	q, ok := divide(10, 2)
	_, ok2 := divide(1, 0) // 不需要的返回值用 _ 丢掉

	if q != 5 || !ok || ok2 {
		t.Error("divide 结果不对")
	}
}

// 可变参数：nums 在函数里就是一个 []int
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func TestVariadic(t *testing.T) {
	s := []int{1, 2, 3}

	if sum(1, 2, 3) != 6 || sum(s...) != 6 || sum() != 0 { // 切片后面加 ... 展开传入
		t.Error("sum 结果不对")
	}
}

// 函数是一等公民：可以赋值给变量、作为参数传递
func apply(f func(int) int, x int) int {
	return f(x)
}

func TestFuncValue(t *testing.T) {
	double := func(x int) int { return x * 2 }

	if apply(double, 3) != 6 {
		t.Error("apply 结果不对")
	}
}

// 闭包：内部函数"捕获"了外部变量 count，每次调用都在改同一个 count
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func TestClosure(t *testing.T) {
	c := counter()
	c()
	c()

	if got := c(); got != 3 {
		t.Errorf("期望 3，实际 %d", got)
	}

	c2 := counter() // 新的闭包，有自己的 count
	if c2() != 1 {
		t.Error("c2 应该从 1 开始")
	}
}

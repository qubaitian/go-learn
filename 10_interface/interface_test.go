package demo

import (
	"fmt"
	"testing"
)

type Shape interface {
	Area() float64
}

type Rect struct{ W, H float64 }
type Circle struct{ R float64 }

// 隐式实现：不需要写 implements，只要方法对得上就算实现了接口
func (r Rect) Area() float64   { return r.W * r.H }
func (c Circle) Area() float64 { return 3 * c.R * c.R }

func totalArea(shapes []Shape) float64 {
	sum := 0.0
	for _, s := range shapes {
		sum += s.Area()
	}
	return sum
}

func TestImplicit(t *testing.T) {
	shapes := []Shape{Rect{2, 3}, Circle{1}}

	if totalArea(shapes) != 9 {
		t.Error("面积不对")
	}
}

// 类型断言：从接口里取出具体类型，用 comma ok 写法就不会 panic
func TestTypeAssertion(t *testing.T) {
	var s Shape = Rect{2, 3}

	r, ok := s.(Rect)
	_, isCircle := s.(Circle)

	if !ok || r.W != 2 || isCircle {
		t.Error("类型断言不对")
	}
}

// type switch：按具体类型分别处理；any 就是 interface{}，什么类型都能放
func describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprint("int:", x)
	case string:
		return "string:" + x
	case nil:
		return "nil"
	default:
		return fmt.Sprintf("其他类型 %T", x)
	}
}

func TestTypeSwitch(t *testing.T) {
	if describe(1) != "int:1" || describe("a") != "string:a" || describe(nil) != "nil" || describe(1.5) != "其他类型 float64" {
		t.Error("type switch 不对")
	}
}

// fmt.Stringer 是标准库里最常见的接口之一：实现了 String()，打印时就会自动调用
func (r Rect) String() string { return fmt.Sprintf("Rect(%vx%v)", r.W, r.H) }

func TestStringer(t *testing.T) {
	if fmt.Sprint(Rect{2, 3}) != "Rect(2x3)" {
		t.Error("Stringer 没有生效")
	}
}

// 大坑：一个"装着 nil 指针"的接口，它本身并不是 nil
// 接口值 = (类型, 值)，只有类型和值都是 nil，接口才等于 nil
type MyErr struct{}

func (*MyErr) Error() string { return "my error" }

func mayFail() error {
	var p *MyErr = nil
	return p // 返回的 error 是 (*MyErr, nil)，不是 nil！
}

func TestNilInterfaceTrap(t *testing.T) {
	err := mayFail()

	t.Logf("err == nil ? %v  动态类型 %T", err == nil, err)

	if err == nil {
		t.Error("坑点演示：这里 err 应该不等于 nil")
	}
	// 正确做法：没有错误时直接 return nil，不要返回一个具体类型的 nil 指针
}

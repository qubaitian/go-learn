package demo

import "testing"

type Counter struct{ n int }

// 值接收者：拿到的是副本，改了也不会影响原值
func (c Counter) IncByValue() { c.n++ }

// 指针接收者：可以修改原值
func (c *Counter) Inc() { c.n++ }

// 只读的方法用值接收者就够了
func (c Counter) Value() int { return c.n }

func TestReceiver(t *testing.T) {
	c := Counter{}
	c.IncByValue()
	c.Inc() // c 是变量，Go 会自动帮你取地址，等价于 (&c).Inc()

	if c.Value() != 1 {
		t.Errorf("期望 1，实际 %d", c.Value())
	}
	// 经验法则：只要有一个方法需要修改接收者，就统一都用指针接收者
}

// 方法不只是 struct 才能有：任何自己定义的类型都可以加方法
type Celsius float64

func (c Celsius) ToFahrenheit() float64 { return float64(c)*9/5 + 32 }

func TestMethodOnType(t *testing.T) {
	if Celsius(100).ToFahrenheit() != 212 {
		t.Error("温度换算不对")
	}
}

// 方法值：把方法绑定到某个对象上，当成普通函数来用
func TestMethodValue(t *testing.T) {
	c := &Counter{}
	inc := c.Inc // 绑定到 c
	inc()
	inc()

	if c.n != 2 {
		t.Errorf("期望 2，实际 %d", c.n)
	}
}

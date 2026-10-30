package demo

import (
	"encoding/json"
	"testing"
)

type Point struct {
	X, Y int
}

// 创建 struct 的几种方式
func TestStructLiteral(t *testing.T) {
	p1 := Point{1, 2}        // 按顺序（不推荐，加字段后容易出错）
	p2 := Point{X: 1, Y: 2}  // 按字段名（推荐）
	p3 := Point{X: 1}        // 没写的字段是零值
	p4 := &Point{X: 1, Y: 2} // 直接拿到指针

	// 指针访问字段不用写 (*p4).X，直接 p4.X 就行
	if p1 != p2 || p3.Y != 0 || p4.X != 1 { // 所有字段都可比较时，struct 可以直接用 == 比较
		t.Error("struct 字面量不对")
	}
}

// struct 是值类型：赋值会复制一份
func TestStructIsValue(t *testing.T) {
	a := Point{X: 1}
	b := a
	b.X = 100

	if a.X != 1 {
		t.Error("修改 b 不应该影响 a")
	}
}

// 匿名 struct：临时用一下，不用专门定义类型
func TestAnonymousStruct(t *testing.T) {
	cfg := struct {
		Host string
		Port int
	}{"localhost", 8080}

	if cfg.Port != 8080 {
		t.Error("匿名 struct 不对")
	}
}

// 嵌入（embedding）：Go 用组合代替继承，被嵌入类型的字段和方法会"提升"上来
type Animal struct{ Name string }

func (a Animal) Hello() string { return "I am " + a.Name }

type Dog struct {
	Animal // 只写类型、不写字段名，就是嵌入
	Breed  string
}

func TestEmbedding(t *testing.T) {
	d := Dog{Animal: Animal{Name: "Lucky"}, Breed: "Corgi"}

	// d.Name 和 d.Hello() 都是从 Animal 提升上来的
	if d.Name != "Lucky" || d.Hello() != "I am Lucky" {
		t.Error("嵌入字段 / 方法不对")
	}
}

// struct tag + 首字母大小写：只有大写开头的字段才能被导出（被 json 等其他包看到）
type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"`
	password string // 小写开头，json 包看不到
}

func TestStructTag(t *testing.T) {
	u := User{Name: "tom", password: "secret"}
	b, _ := json.Marshal(u)

	t.Log(string(b))

	if string(b) != `{"name":"tom"}` { // Age 是 0，被 omitempty 省略了
		t.Errorf("json 结果不对: %s", b)
	}
}

package demo

import (
	"maps"
	"slices"
	"testing"
)

// 增删改查
func TestMapBasic(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	m["c"] = 3     // 增 / 改
	delete(m, "a") // 删

	if len(m) != 2 || m["c"] != 3 {
		t.Error("map 操作结果不对", m)
	}
}

// 坑：读一个不存在的 key 不会报错，而是返回零值；要用 "comma ok" 区分
func TestCommaOk(t *testing.T) {
	m := map[string]int{"zero": 0}

	v1, ok1 := m["zero"] // 存在，值是 0
	v2, ok2 := m["nope"] // 不存在，也返回 0

	t.Log(v1, ok1, v2, ok2)

	if !ok1 || ok2 {
		t.Error("comma ok 结果不对")
	}
}

// 坑：遍历 map 的顺序是随机的；需要固定顺序就先把 key 排序
func TestIterateSorted(t *testing.T) {
	m := map[string]int{"b": 2, "a": 1, "c": 3}

	keys := slices.Sorted(maps.Keys(m)) // Go 1.23+

	if !slices.Equal(keys, []string{"a", "b", "c"}) {
		t.Error("排序后的 key 不对", keys)
	}
}

// 坑：nil map 可以读，但写入会 panic，一定要先 make
func TestNilMap(t *testing.T) {
	var m map[string]int
	_ = m["x"] // 读没问题

	defer func() {
		if r := recover(); r == nil { // recover 的用法见 12_defer_panic
			t.Error("向 nil map 写入应该 panic")
		} else {
			t.Log("果然 panic 了:", r)
		}
	}()
	m["x"] = 1
}

// map 是引用语义：传给函数后，函数里的修改调用方能看到
func addOne(m map[string]int) { m["new"] = 1 }

func TestMapIsReference(t *testing.T) {
	m := make(map[string]int)
	addOne(m)

	if m["new"] != 1 {
		t.Error("函数里的修改应该可见")
	}
}

// 坑：map 里存的是 struct 值时，不能直接改它的字段
func TestMapOfStruct(t *testing.T) {
	type User struct{ Age int }

	m := map[string]User{"tom": {Age: 1}}
	// m["tom"].Age = 2  // 编译错误：cannot assign to struct field in map
	u := m["tom"] // 办法 1：取出来，改完再放回去
	u.Age = 2
	m["tom"] = u

	pm := map[string]*User{"tom": {Age: 1}} // 办法 2：存指针
	pm["tom"].Age = 2

	if m["tom"].Age != 2 || pm["tom"].Age != 2 {
		t.Error("修改失败")
	}
}

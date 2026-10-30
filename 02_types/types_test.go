package demo

import (
	"strconv"
	"strings"
	"testing"
)

// Go 没有隐式类型转换，int 和 float64 不能直接相加
func TestConversion(t *testing.T) {
	i := 3
	f := 1.5
	sum := float64(i) + f // 必须手动转换，写 i + f 编译不过

	if sum != 4.5 {
		t.Errorf("期望 4.5，实际 %v", sum)
	}
}

// 坑：数字转字符串要用 strconv，string(65) 得到的是 "A" 而不是 "65"
func TestNumberToString(t *testing.T) {
	s1 := strconv.Itoa(65)
	s2 := string(rune(65))

	t.Logf("strconv.Itoa(65)=%q  string(rune(65))=%q", s1, s2)

	if s1 != "65" || s2 != "A" {
		t.Error("转换结果不符合预期")
	}

	n, err := strconv.Atoi("123") // 字符串转数字，可能失败，所以有 err
	if err != nil || n != 123 {
		t.Errorf("Atoi 失败: %v", err)
	}
}

// string 底层是 UTF-8 字节；len 返回的是字节数，不是字符数
func TestStringBytesRunes(t *testing.T) {
	s := "Go你好"

	t.Log("len(s) =", len(s), "字符数 =", len([]rune(s)))

	if len(s) != 8 { // G o 各 1 字节，每个汉字 3 字节
		t.Errorf("期望 8 字节，实际 %d", len(s))
	}

	// range 遍历 string 得到的是 rune（字符），i 是字节下标
	for i, r := range s {
		t.Logf("下标 %d: %c", i, r)
	}

	// 下标访问得到的是 byte，不是字符
	if s[0] != 'G' {
		t.Error("s[0] 应该是 'G'")
	}
}

// string 不可变；大量拼接用 strings.Builder
func TestStringBuilder(t *testing.T) {
	var b strings.Builder
	for i := range 3 {
		b.WriteString(strconv.Itoa(i))
	}

	if b.String() != "012" {
		t.Errorf("期望 012，实际 %s", b.String())
	}
}

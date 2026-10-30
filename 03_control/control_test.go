package demo

import (
	"strconv"
	"testing"
)

// if 可以带一个初始化语句，变量只在 if/else 里有效
func TestIfInit(t *testing.T) {
	if n, err := strconv.Atoi("42"); err == nil {
		t.Log("解析成功:", n)
	} else {
		t.Error("不应该失败", err)
	}
}

// Go 只有 for，没有 while
func TestFor(t *testing.T) {
	sum := 0
	for i := 0; i < 5; i++ { // 经典三段式
		sum += i
	}

	n := 0
	for n < 3 { // 相当于 while
		n++
	}

	for i := range 3 { // Go 1.22+：直接 range 一个整数，i = 0,1,2
		t.Log("range int:", i)
	}

	if sum != 10 || n != 3 {
		t.Errorf("sum=%d n=%d", sum, n)
	}
}

// switch 默认不会往下穿透，不用写 break
func TestSwitch(t *testing.T) {
	grade := func(score int) string {
		switch { // 不写表达式，相当于 switch true，可以替代 if-else 链
		case score >= 90:
			return "A"
		case score >= 60:
			return "B"
		default:
			return "C"
		}
	}

	day := "sat"
	weekend := false
	switch day {
	case "sat", "sun": // 一个 case 可以匹配多个值
		weekend = true
	}

	if grade(95) != "A" || grade(70) != "B" || grade(10) != "C" || !weekend {
		t.Error("switch 结果不对")
	}
}

// 带标签的 break 可以一次跳出多层循环
func TestLabeledBreak(t *testing.T) {
	found := ""
outer:
	for i := range 3 {
		for j := range 3 {
			if i*j == 2 {
				found = strconv.Itoa(i) + "," + strconv.Itoa(j)
				break outer
			}
		}
	}

	if found != "1,2" {
		t.Errorf("期望 1,2，实际 %s", found)
	}
}

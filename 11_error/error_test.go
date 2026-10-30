package demo

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

// Go 没有异常，错误就是一个普通的返回值：error 接口只有一个 Error() string 方法
func parseAge(s string) (int, error) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parseAge(%q): %w", s, err) // %w 把原错误"包装"进来
	}
	if n < 0 {
		return 0, errors.New("年龄不能是负数")
	}
	return n, nil
}

func TestBasicError(t *testing.T) {
	if _, err := parseAge("18"); err != nil {
		t.Error("不应该出错", err)
	}

	_, err := parseAge("abc")
	t.Log(err)
	if err == nil {
		t.Error("应该出错")
	}
}

// 哨兵错误（sentinel error）+ errors.Is：判断错误链里有没有某个特定的错误
var ErrNotFound = errors.New("not found")

func findUser(id int) error {
	return fmt.Errorf("findUser %d: %w", id, ErrNotFound)
}

func TestErrorsIs(t *testing.T) {
	err := findUser(1)

	if err == ErrNotFound { // 被包装过了，所以 == 判断不出来
		t.Error("包装后不应该 ==")
	}
	if !errors.Is(err, ErrNotFound) { // errors.Is 会沿着包装链一层层往下找
		t.Error("errors.Is 应该找到 ErrNotFound")
	}
}

// 自定义错误类型 + errors.As：从错误链里取出具体的错误类型，读它的字段
type ValidationError struct {
	Field string
}

func (e *ValidationError) Error() string { return e.Field + " 不合法" }

func validate() error {
	return fmt.Errorf("validate: %w", &ValidationError{Field: "email"})
}

func TestErrorsAs(t *testing.T) {
	err := validate()

	var ve *ValidationError
	if !errors.As(err, &ve) || ve.Field != "email" {
		t.Error("errors.As 应该取出 ValidationError")
	}
}

// errors.Join：把多个错误合成一个（Go 1.20+）
func TestErrorsJoin(t *testing.T) {
	err := errors.Join(ErrNotFound, errors.New("另一个错误"))

	t.Log(err)

	if !errors.Is(err, ErrNotFound) {
		t.Error("Join 之后也能用 errors.Is 找到")
	}
}

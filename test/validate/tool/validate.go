package tool

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

//=================== 定义接口 ==========================

// Validator 验证接口
type Validator interface {
	Validate(interface{}) (bool, error)
}

// DefaultValidator 默认
type DefaultValidator struct{}

// Validate 默认
func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

// =================== 定义接口 end ==========================

//==================== 校验 ======================
const tagName = "validate"

// ValidateStruct 校验结构体
func ValidateStruct(s interface{}) []error {
	var errs []error
	v := reflect.ValueOf(s)

	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagName)

		if tag == "" || tag == "-" {
			continue
		}

		validator := getValidatorFromTag(tag)

		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}

func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "number":
		validator := NumberValidator{}
		//将structTag中的min和max解析到结构体中
		_, _ = fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "string":
		validator := StringValidator{}
		_, _ = fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "email":
		return EmailValidator{}
	}

	return DefaultValidator{}
}

//==================== 校验 end ======================

// 邮箱校验规则 ==============================

// EmailValidator ....
type EmailValidator struct{}

// 邮箱验证正则
var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

// Validate ...
func (v EmailValidator) Validate(val interface{}) (bool, error) {
	if !mailRe.MatchString(val.(string)) {
		return false, fmt.Errorf("邮箱格式不正确")
	}
	return true, nil
}

// 邮箱校验规则 =================end=============

// 字符串校验规则 =====================

// StringValidator ...
type StringValidator struct {
	Min int
	Max int
}

// Validate ...
func (v StringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))

	if l == 0 {
		return false, fmt.Errorf("字符串不能为空")
	}

	if l < v.Min {
		return false, fmt.Errorf("字符长度必须大于 %v", v.Min)
	}

	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf("字符长度必须小于 %v", v.Max)
	}

	return true, nil
}

// 字符串校验规则 end=====================

// 数字校验规则 =========================

// NumberValidator ...
type NumberValidator struct {
	Min int
	Max int
}

// Validate ...
func (v NumberValidator) Validate(val interface{}) (bool, error) {
	num := val.(int)

	if num < v.Min {
		return false, fmt.Errorf("should be greater than %v", v.Min)
	}

	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("should be less than %v", v.Max)
	}

	return true, nil
}

// 数字校验规则 =========================

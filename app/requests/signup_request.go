// Package requests 处理请求数据和表单验证
package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

// SignupPhoneExist 验证注册手机号是否存在表单
func SignupPhoneExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义校验规则
	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	// 自定义验证出错时的提示
	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号码为必填，参数名称 phone",
			"digits:手机号码长度必须为 11 位的数字",
		},
	}

	return validate(data, rules, messages)
}

// SignupEmailExist 验证注册邮箱是否存在表单
func SignupEmailExist(data interface{}, c *gin.Context) map[string][]string {

	// 自定义规则
	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	// 自定义验证失败错误提示
	messages := govalidator.MapData{
		"email": []string{
			"required:Email 为必填项",
			"min:Email 长度最小为 4",
			"max:Email 长度最大为 30",
			"email:Email 格式不正确，请提供正确的 Email 格式",
		},
	}

	return validate(data, rules, messages)
}

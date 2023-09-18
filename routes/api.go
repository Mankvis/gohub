package routes

import (
	"github.com/gin-gonic/gin"
	"gohub/app/http/controllers/api/v1/auth"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有的 v1 版本的路由都将存在这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已经存在
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断邮箱是否已经存在
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)
			// 手机+验证码注册
			authGroup.POST("/signup/using-phone", suc.SignupUsingPhone)
			// 邮箱+验证码注册
			authGroup.POST("/signup/using-email", suc.SignupUsingEmail)
			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
			authGroup.POST("/verify-codes/phone", vcc.SendUsingPhone)
			authGroup.POST("/verify-codes/email", vcc.SendUsingEmail)
			// 登录
			lgc := new(auth.LoginController)
			// 手机号+短信
			authGroup.POST("/login/using-phone", lgc.LoginByPhone)
			// 手机号、email 和用户名 + 密码
			authGroup.POST("/login/using-password", lgc.LoginByPassword)
			// 刷新 token
			authGroup.POST("/login/refresh-token", lgc.RefreshToken)
			// 重置密码
			pwc := new(auth.PasswordController)
			authGroup.POST("/password-reset/using-phone", pwc.ResetByPhone)
			authGroup.POST("/password-reset/using-email", pwc.ResetByEmail)
		}
	}
}

package api

import (
	"github.com/20326/vega/app/handler/api/user"
	"github.com/gin-gonic/gin"
)

func NewHandlers(r *gin.Engine) {
	// api
	apiGroup := r.Group("/api")

	apiGroup.POST("/user/auth/2step-code", user.Step2CodeAction)
	apiGroup.POST("/user/register", user.RegisterAction)
	apiGroup.POST("/user/login", user.LoginAction)
	apiGroup.POST("/user/logout", user.LogoutAction)
	apiGroup.POST("/user/change-password", user.ChangePasswordAction)
	apiGroup.POST("/user/forget-password", user.ForgetPasswordAction)
	apiGroup.POST("/user/reset-password", user.ResetPasswordAction)
	apiGroup.GET("/user/info", user.UserInfoAction)
}

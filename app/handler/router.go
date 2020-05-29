package handler

import (
	// "net/http"

	"github.com/20326/vega/app/handler/api/action"
	"github.com/20326/vega/app/handler/api/permission"
	"github.com/20326/vega/app/handler/api/role"
	"github.com/20326/vega/app/handler/api/setting"
	"github.com/20326/vega/app/handler/api/user"
	"github.com/gin-gonic/gin"
	// "github.com/phuslu/log"
)

func NewHandlers(r *gin.Engine) {

	// api
	apiGroup := r.Group("/api")
	{
		// user action login/logout
		apiGroup.GET("/user/auth/test-code", user.TestAction)
		apiGroup.POST("/user/auth/2step-code", user.Step2CodeAction)
		apiGroup.POST("/user/register", user.RegisterAction)
		apiGroup.POST("/user/login", user.LoginAction)
		apiGroup.POST("/user/logout", user.LogoutAction)
		apiGroup.POST("/user/change-password", user.ChangePasswordAction)
		apiGroup.POST("/user/forget-password", user.ForgetPasswordAction)
		apiGroup.POST("/user/reset-password", user.ResetPasswordAction)
		apiGroup.GET("/user/info", user.UserInfoAction)

		apiGroup.GET("/settings", setting.GetSettingsAction)
		apiGroup.GET("/settings/:id", setting.GetSettingAction)
		// apiGroup.POST("/settings", setting.UpdateSettingsAction)
		apiGroup.PUT("/settings/:id", setting.UpdateSettingAction)
		apiGroup.DELETE("/settings/:id", setting.DeleteSettingAction)

		// action admin
		apiGroup.GET("/actions", action.GetActionsAction)
		apiGroup.POST("/actions", action.AddActionAction)
		apiGroup.GET("/actions/:id", action.GetActionAction)
		apiGroup.PUT("/actions/:id", action.UpdateActionAction)
		apiGroup.DELETE("/actions/:id", action.DeleteActionAction)
		// role admin
		apiGroup.GET("/roles", role.GetRolesAction)
		apiGroup.POST("/roles", role.AddRoleAction)
		apiGroup.GET("/roles/:id", role.GetRoleAction)
		apiGroup.PUT("/roles/:id", role.UpdateRoleAction)
		apiGroup.DELETE("/roles/:id", role.DeleteRoleAction)
		// permission admin
		apiGroup.GET("/permissions", permission.GetPermissionAction)
		apiGroup.POST("/permissions", permission.AddPermissionAction)
		apiGroup.GET("/permissions/:id", permission.GetPermissionAction)
		apiGroup.PUT("/permissions/:id", permission.UpdatePermissionAction)
		apiGroup.DELETE("/permissions/:id", permission.DeletePermissionAction)
	}

	// console
	//consoleGroup := r.Group("/api/console")
	//{
	//	// user
	//	consoleGroup.GET("/users", console.GetUsersPageAction)
	//	consoleGroup.POST("/users", console.AddUserAction)
	//	consoleGroup.GET("/users/:id", console.GetUserAction)
	//	consoleGroup.PUT("/users/:id", console.UpdateUserAction)
	//	consoleGroup.DELETE("/users/:id", console.DeleteUserAction)
	//	// role admin
	//	consoleGroup.GET("/roles", console.GetRolesAction)
	//	consoleGroup.POST("/roles", console.AddRoleAction)
	//	consoleGroup.GET("/roles/:id", console.GetRoleAction)
	//	consoleGroup.PUT("/roles/:id", console.UpdateRoleAction)
	//	consoleGroup.DELETE("/roles/:id", console.DeleteRoleAction)
	//	// permission admin
	//	consoleGroup.GET("/permissions", console.GetPermissionAction)
	//	consoleGroup.POST("/permissions", console.AddPermissionAction)
	//	consoleGroup.GET("/permissions/:id", console.GetPermissionAction)
	//	consoleGroup.PUT("/permissions/:id", console.UpdatePermissionAction)
	//	consoleGroup.DELETE("/permissions/:id", console.DeletePermissionAction)
	//	// setting admin
	//	consoleGroup.GET("/settings", console.GetSettingsAction)
	//	consoleGroup.GET("/settings/:category", console.GetSettingsAction)
	//	consoleGroup.POST("/settings", console.UpdateSettingsAction)
	//	consoleGroup.PUT("/settings/:id", console.UpdateSettingAction)
	//	consoleGroup.DELETE("/settings/:id", console.DeleteSettingAction)
	//	// user account
	//	consoleGroup.GET("/user/profile", console.GetUserProfileAction)
	//	consoleGroup.POST("/user/profile", console.UpdateUserProfileAction)
	//
	//}

	// portal

	// theme

	// static

	// casbin
	//{
	//
	//	r.GET("/init", func(c *gin.Context) {
	//		service.Casbin.Enforcer().AddPolicy("role_admin", "/data", "read")
	//		service.Casbin.Enforcer().AddPolicy("role_admin", "/book/:id", "read")
	//		service.Casbin.Enforcer().AddGroupingPolicy("admin", "role_admin")
	//		service.Casbin.Enforcer().AddGroupingPolicy("test", "role_admin")
	//
	//		service.Casbin.Enforcer().SavePolicy()
	//		c.String(http.StatusOK, "Add Policy Finised")
	//	})
	//
	//	r.GET("/", func(c *gin.Context) {
	//		// test begin
	//		log.Info().Msgf("users: %v", service.User.GetUsers())
	//
	//		hasPermission, err := service.Casbin.CheckPermissionission("admin", "/data", "read")
	//		log.Info().Err(err).Msgf("admin perm: /data read %v", hasPermission)
	//		hasPermission, err = service.Casbin.CheckPermissionission("admin", "/book/121", "view")
	//		log.Info().Err(err).Msgf("admin perm: /book/121 view %v", hasPermission)
	//		hasPermission, err = service.Casbin.CheckPermissionission("admin", "/book/121", "read")
	//		log.Info().Err(err).Msgf("admin perm: /book/121 read %v", hasPermission)
	//
	//		hasPermission, err = service.Casbin.CheckPermissionission("test", "/book/121", "read")
	//		log.Info().Err(err).Msgf("test perm: /book/121 read %v", hasPermission)
	//		hasPermission, err = service.Casbin.CheckPermissionission("test", "/book/121", "view")
	//		log.Info().Err(err).Msgf("test perm: /book/121 view %v", hasPermission)
	//		hasPermission, err = service.Casbin.CheckPermissionission("test", "/data", "read")
	//		log.Info().Err(err).Msgf("test perm: /data read %v", hasPermission)
	//
	//		c.String(http.StatusOK, "Welcome Gin Http Server")
	//	})
	//}
}

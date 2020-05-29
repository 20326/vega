package console

import (
	"github.com/20326/vega/app/handler/console/action"
	"github.com/20326/vega/app/handler/console/permission"
	"github.com/20326/vega/app/handler/console/role"
	"github.com/20326/vega/app/handler/console/setting"
	"github.com/20326/vega/app/handler/console/user"
	"github.com/gin-gonic/gin"
)

func NewHandlers(r *gin.Engine) {

	// console
	consoleGroup := r.Group("/api/console")

	consoleGroup.GET("/actions", action.GetActionsAction)
	consoleGroup.POST("/actions", action.AddActionAction)
	consoleGroup.GET("/actions/:id", action.GetActionAction)
	consoleGroup.PUT("/actions/:id", action.UpdateActionAction)
	consoleGroup.DELETE("/actions/:id", action.DeleteActionAction)

	consoleGroup.GET("/roles", role.GetRolesAction)
	consoleGroup.POST("/roles", role.AddRoleAction)
	consoleGroup.GET("/roles/:id", role.GetRoleAction)
	consoleGroup.PUT("/roles/:id", role.UpdateRoleAction)
	consoleGroup.DELETE("/roles/:id", role.DeleteRoleAction)

	consoleGroup.GET("/settings", setting.GetSettingsAction)
	consoleGroup.GET("/settings/:id", setting.GetSettingAction)
	// consoleGroup.POST("/settings", setting.UpdateSettingsAction)
	//consoleGroup.GET("/settings/:category", setting.GetSettingsAction)
	consoleGroup.PUT("/settings/:id", setting.UpdateSettingAction)
	consoleGroup.DELETE("/settings/:id", setting.DeleteSettingAction)

	consoleGroup.GET("/permissions", permission.GetPermissionAction)
	consoleGroup.POST("/permissions", permission.AddPermissionAction)
	consoleGroup.GET("/permissions/:id", permission.GetPermissionAction)
	consoleGroup.PUT("/permissions/:id", permission.UpdatePermissionAction)
	consoleGroup.DELETE("/permissions/:id", permission.DeletePermissionAction)

	consoleGroup.GET("/users", user.GetUserAction)
	consoleGroup.POST("/users", user.AddUserAction)
	consoleGroup.GET("/users/:id", user.GetUserAction)
	consoleGroup.PUT("/users/:id", user.UpdateUserAction)
	consoleGroup.DELETE("/users/:id", user.DeleteUserAction)
}

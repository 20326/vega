package console

import (
	"errors"
	"net/http"

	"github.com/20326/vega/app/handler/console/action"
	"github.com/20326/vega/app/handler/console/permission"
	"github.com/20326/vega/app/handler/console/role"
	"github.com/20326/vega/app/handler/console/setting"
	"github.com/20326/vega/app/handler/console/user"
	"github.com/20326/vega/pkg/render"
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
	consoleGroup.GET("/settings/:id", GetSettingsHandler)
	consoleGroup.POST("/settings/:id", UpdateSettingsHandler)
	consoleGroup.GET("/settings/:id/:group", GetSettingsHandler)
	consoleGroup.POST("/settings/:id/:group", UpdateSettingsHandler)
	consoleGroup.PUT("/settings/:id", setting.AddSettingAction)
	consoleGroup.DELETE("/settings/:id", setting.DeleteSettingAction)

	consoleGroup.GET("/permissions", permission.GetPermissionsAction)
	consoleGroup.POST("/permissions", permission.AddPermissionAction)
	consoleGroup.GET("/permissions/:id", permission.GetPermissionAction)
	consoleGroup.PUT("/permissions/:id", permission.UpdatePermissionAction)
	consoleGroup.DELETE("/permissions/:id", permission.DeletePermissionAction)

	consoleGroup.GET("/users", user.GetUsersAction)
	consoleGroup.POST("/users", user.AddUserAction)
	consoleGroup.GET("/users/:id", user.GetUserAction)
	consoleGroup.PUT("/users/:id", user.UpdateUserAction)
	consoleGroup.DELETE("/users/:id", user.DeleteUserAction)
	consoleGroup.GET("/user/profile", user.GetCurrentUserAction)
	consoleGroup.POST("/user/profile", user.UpdateCurrentUserAction)
}

// TODO uglify code
func GetSettingsHandler(c *gin.Context) {
	idArg := c.Param("id")
	groupArg := c.Param("group")

	if "group" == idArg && "" != groupArg{
		setting.GetSettingsAction(c)
	} else if "" == groupArg {
		setting.GetSettingAction(c)
	} else {
		result := render.NewResult()
		result.Error(errors.New("not found path"))
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
}

// TODO uglify code
func UpdateSettingsHandler(c *gin.Context) {
	idArg := c.Param("id")
	groupArg := c.Param("group")

	if "group" == idArg && "" != groupArg{
		setting.UpdateSettingsAction(c)
	} else if "" == groupArg {
		setting.UpdateSettingAction(c)
	} else {
		result := render.NewResult()
		result.Error(errors.New("not found path"))
		c.AbortWithStatusJSON(http.StatusInternalServerError, result)
	}
}
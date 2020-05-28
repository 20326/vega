package perm
//
//import (
//	"net/http"
//	"strconv"
//
//	"github.com/20326/vega/app/model"
//	"github.com/20326/vega/app/service"
//	"github.com/20326/vega/pkg/render"
//	// "github.com/phuslu/log"
//	"github.com/gin-gonic/gin"
//)
//
//// GetPermissionsAction gets permissions.
//func GetPermissionsAction(c *gin.Context) {
//	result := render.NewResult()
//	defer c.JSON(http.StatusOK, result)
//
//	permissionModels := service.Permission.GetPermissions()
//
//	var permissions []*model.ConsolePermission
//	for _, permissionModel := range permissionModels {
//		comment := &model.ConsolePermission{
//			ID:   permissionModel.ID,
//			Name: permissionModel.Name,
//		}
//
//		permissions = append(permissions, comment)
//	}
//
//	data := map[string]interface{}{}
//	data["permissions"] = permissions
//	result.Result = data
//}
//
//// GetPermissionAction get a permission.
//func GetPermissionAction(c *gin.Context) {
//	result := render.NewResult()
//	defer c.JSON(http.StatusOK, result)
//
//	idArg := c.Param("id")
//	id, err := strconv.ParseUint(idArg, 10, 64)
//	if nil != err {
//		result.Code = pkg.CodeErr
//
//		return
//	}
//
//	data := service.Permission.GetPermission(id)
//	if nil == data {
//		result.Code = pkg.CodeErr
//
//		return
//	}
//
//	result.Result = data
//}
//
//// DeletePermissionAction remove a permission.
//func DeletePermissionAction(c *gin.Context) {
//	result := render.NewResult()
//	defer c.JSON(http.StatusOK, result)
//
//	idArg := c.Param("id")
//	id, err := strconv.ParseUint(idArg, 10, 64)
//	if nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = err.Error()
//
//		return
//	}
//
//	if err := service.Permission.DeletePermission(id); nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = err.Error()
//	}
//}
//
//// UpdatePermissionAction updates a permission.
//func UpdatePermissionAction(c *gin.Context) {
//	result := render.NewResult()
//	defer c.JSON(http.StatusOK, result)
//
//	idArg := c.Param("id")
//	id, err := strconv.ParseUint(idArg, 10, 64)
//	if nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = err.Error()
//
//		return
//	}
//
//	permission := &model.Permission{Model: model.Model{ID: uint64(id)}}
//	if err := c.BindJSON(permission); nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = "parses update permission request failed"
//
//		return
//	}
//
//	if err := service.Permission.UpdatePermission(permission); nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = err.Error()
//	}
//}
//
//// AddPermissionAction adds a permission.
//func AddPermissionAction(c *gin.Context) {
//	result := render.NewResult()
//	defer c.JSON(http.StatusOK, result)
//
//	permission := &model.Permission{}
//	if err := c.BindJSON(permission); nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = "parses add permission request failed"
//
//		return
//	}
//
//	if err := service.Permission.AddPermission(permission); nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = err.Error()
//	}
//}

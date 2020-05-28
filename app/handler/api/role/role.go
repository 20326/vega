package role
//
//import (
//	"net/http"
//	"strconv"
//
//	"github.com/20326/vega/app/model"
//	"github.com/20326/vega/app/service"
//	"github.com/20326/vega/pkg"
//	// "github.com/phuslu/log"
//	"github.com/gin-gonic/gin"
//)
//
//// GetRolesAction gets roles.
//func GetRolesAction(c *gin.Context) {
//	result := render.NewResult()
//	defer c.JSON(http.StatusOK, result)
//
//	roleModels := service.Role.GetRoles()
//
//	var roles []*model.ConsoleRole
//	for _, roleModel := range roleModels {
//		comment := &model.ConsoleRole{
//			ID:   roleModel.ID,
//			Name: roleModel.Name,
//		}
//
//		roles = append(roles, comment)
//	}
//
//	data := map[string]interface{}{}
//	data["roles"] = roles
//	result.Result = data
//}
//
//// GetRoleAction get a role.
//func GetRoleAction(c *gin.Context) {
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
//	data := service.Role.ConsoleGetRole(id)
//	if nil == data {
//		result.Code = pkg.CodeErr
//
//		return
//	}
//
//	result.Result = data
//}
//
//// DeleteRoleAction remove a role.
//func DeleteRoleAction(c *gin.Context) {
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
//	if err := service.Role.DeleteRole(id); nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = err.Error()
//	}
//}
//
//// UpdateRoleAction updates a role.
//func UpdateRoleAction(c *gin.Context) {
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
//	role := &model.Role{Model: model.Model{ID: uint64(id)}}
//	if err := c.BindJSON(role); nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = "parses update role request failed"
//
//		return
//	}
//
//	if err := service.Role.UpdateRole(role); nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = err.Error()
//	}
//}
//
//// AddRoleAction adds a role.
//func AddRoleAction(c *gin.Context) {
//	result := render.NewResult()
//	defer c.JSON(http.StatusOK, result)
//
//	role := &model.Role{}
//	if err := c.BindJSON(role); nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = "parses add role request failed"
//
//		return
//	}
//
//	if err := service.Role.AddRole(role); nil != err {
//		result.Code = pkg.CodeErr
//		result.Msg = err.Error()
//	}
//}

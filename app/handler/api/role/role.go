package role

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/render"
	"github.com/gin-gonic/gin"
)

// GetRolesAction gets roles.
func GetRolesAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	roles, _ := srv.Roles.List(c)

	//var roles []*model.ConsoleRole
	//for _, roleModel := range roleModels {
	//	comment := &model.ConsoleRole{
	//		ID:   roleModel.ID,
	//		Name: roleModel.Name,
	//	}
	//
	//	roles = append(roles, comment)
	//}

	data := map[string]interface{}{}
	data["roles"] = roles
	result.Result = data
}

// GetRoleAction get a role.
func GetRoleAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	data, err := srv.Roles.Find(c, id)
	if nil == data {
		result.Error(err)

		return
	}

	result.Result = data
}

// DeleteRoleAction remove a role.
func DeleteRoleAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Roles.Delete(c, id); nil != err {
		result.Error(err)

	}
}

// UpdateRoleAction updates a role.
func UpdateRoleAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	role := &model.Role{Model: model.Model{ID: uint64(id)}}
	if err := c.BindJSON(role); nil != err {
		result.Error(errors.New("parses update role request failed"))

		return
	}

	srv := service.FromContext(c)
	if err := srv.Roles.Update(c, role); nil != err {
		result.Error(err)
	}
}

// AddRoleAction adds a role.
func AddRoleAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	role := &model.Role{}
	if err := c.BindJSON(role); nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Roles.Create(c, role); nil != err {
		result.Error(err)
	}
}

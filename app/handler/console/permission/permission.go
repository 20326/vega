package permission

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/params"
	"github.com/20326/vega/pkg/render"
	"github.com/gin-gonic/gin"
)

// GetPermissionsAction gets permissions.
func GetPermissionsAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)

	where := ""
	whereArgs := []interface{}{""}
	// TODO

	pageQuery := model.PageQuery{
		PageNo: params.GetIntArgs(c, "pageNo"),
		PageSize: params.GetIntArgs(c, "pageSize"),
		Where: where,
		WhereArgs: whereArgs,
	}

	permissions, pagination := srv.Permissions.FindWhere(pageQuery)

	pagination.SetData(permissions)
	result.Result = pagination
}

// GetPermissionAction get a permission.
func GetPermissionAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	data, err := srv.Permissions.Find(c, id)
	if nil == data {
		result.Error(err)

		return
	}

	result.Result = data
}

// DeletePermissionAction remove a permission.
func DeletePermissionAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Permissions.Delete(c, id); nil != err {
		result.Error(err)

	}
}

// UpdatePermissionAction updates a permission.
func UpdatePermissionAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	permission := &model.Permission{Model: model.Model{ID: uint64(id)}}
	if err := c.BindJSON(permission); nil != err {
		result.Error(errors.New("parses update permission request failed"))

		return
	}

	srv := service.FromContext(c)
	if err := srv.Permissions.Update(c, permission); nil != err {
		result.Error(err)
	}
}

// AddPermissionAction adds a permission.
func AddPermissionAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	permission := &model.Permission{}
	if err := c.BindJSON(permission); nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Permissions.Create(c, permission); nil != err {
		result.Error(err)
	}
}

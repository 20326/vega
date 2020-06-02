package role

import (
	"errors"
	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/array"
	"github.com/20326/vega/pkg/render"
	mapset "github.com/deckarep/golang-set"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

// GetRolesAction gets roles.
func GetRolesAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	roles, _ := srv.Roles.List(c)

	// TODO get select actions data
	for _, role := range roles{
		roleActionIDs := role.GetActionIds()
		//rolePermissions = append(rolePermissions, &model.RolePermissions{
		//	RoleID: role.ID,
		//	PermissionID: permission.ID,
		//	ActionsData: strings.Join(permission.Selected, ","),
		//})
		for _, permission := range role.Permissions{
			permActionIDs := permission.GetActionIds()
			permission.Selected = roleActionIDs.Intersect(permActionIDs).ToSlice()
		}
		//if permission.Actions
	}
	data := map[string]interface{}{}
	data["data"] = roles
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
	log := srv.GetLogger()

	log.WithFields(logrus.Fields{
		"action":   "UpdateRoleAction",
		"role.Permissions": role.Permissions,
	}).Info("update role permissions")

	var actionIds = mapset.NewSet()
	for _, permission := range role.Permissions{
		//rolePermissions = append(rolePermissions, &model.RolePermissions{
		//	RoleID: role.ID,
		//	PermissionID: permission.ID,
		//	ActionsData: strings.Join(permission.Selected, ","),
		//})

		for _, action := range permission.Actions{
			logrus.Printf("%d  %s %+v %d\n", action.ID, action.Name, permission.Selected, array.IndexOf(action.ID, permission.Selected))
			if 0< len(permission.Selected) {
				actionIds = actionIds.Union(mapset.NewSetFromSlice(permission.Selected))
			}
			// actions = append(actions, action)
		}
		//if permission.Actions
	}
	logrus.Printf("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n %+v \n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n", actionIds)
	role.Actions = nil
	//role, err := srv.Roles.Find(c, arg["id"].(uint64))
	//if nil != err {
	//	result.Error(errors.New("not found role, failed"))
	//
	//	return
	//}

	// srv.Roles.RelatedClear(c, role)

	if err := srv.Roles.Update(c, role, actionIds.ToSlice()); nil != err {
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

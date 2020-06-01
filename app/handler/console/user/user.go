package user

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/crypto"
	"github.com/20326/vega/pkg/params"
	"github.com/20326/vega/pkg/render"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// GetUsersAction gets users.
func GetUsersAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)

	where := []string{}
	whereArgs := []interface{}{}

	name := c.Query("name")
	if "" != name {
		where = append(where, "(`username` LIKE ? OR `nickname` LIKE ?)")
		whereArgs = append(whereArgs, "%"+name+"%", "%"+name+"%")
	}

	status := params.GetIntArgs(c, "status")
	if 0 < status {
		where = append(where, "`status` = ?")
		whereArgs = append(whereArgs, status)
	}

	phone := c.Query("phone")
	if 0 < len(phone) {
		where = append(where, "`phone` LIKE ?")
		whereArgs = append(whereArgs, "%"+phone+"%")
	}

	roles := c.QueryArray("role[]")

	pageQuery := model.PageQuery{
		PageNo:    params.GetIntArgs(c, "pageNo"),
		PageSize:  params.GetIntArgs(c, "pageSize"),
		Where:     strings.Join(where, " AND "),
		WhereArgs: whereArgs,
	}
	users, pagination := srv.Users.FindWhere(pageQuery, roles)

	pagination.SetData(users)
	result.Result = pagination
}

// GetUserAction get a user.
func GetUserAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	data, err := srv.Users.Find(c, id)
	if nil == data {
		result.Error(err)

		return
	}

	result.Result = data
}

// DeleteUserAction remove a user.
func DeleteUserAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Users.Delete(c, id); nil != err {
		result.Error(err)

	}
}

// UpdateUserAction updates a user.
func UpdateUserAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	user := &model.User{Model: model.Model{ID: uint64(id)}}
	if err := c.BindJSON(user); nil != err {
		result.Error(errors.New("parses update user request failed"))

		return
	}

	srv := service.FromContext(c)
	oldUser, err := srv.Users.Find(c, user.ID)
	if nil == oldUser {
		result.Error(err)

		return
	}
	// TODO, 检查Username Nickname 重复, unique_index
	oldUser.Username = user.Username
	oldUser.Nickname = user.Nickname
	oldUser.Phone = user.Phone
	oldUser.Email = user.Email
	oldUser.UpdatedAt = time.Now()
	oldUser.BIO = user.BIO
	oldUser.Status = user.Status

	if 5 <= len(user.Password) {
		oldUser.PasswordHash = crypto.HashAndSalt([]byte(user.Password))
		oldUser.Token = uuid.NewV4().String()
	}
	oldUser.Password = ""

	// 清除用户角色相关性
	roles := []*model.Role{}
	for _, roleID := range user.RoleList {
		role, _ := srv.Roles.Find(c, roleID)
		if nil != role {
			roles = append(roles, role)
		}
	}
	if 0 < len(roles) {
		srv.Users.RelatedClear(c, oldUser)
		oldUser.Roles = roles
	}

	if err := srv.Users.Update(c, oldUser); nil != err {
		result.Error(err)
	}
}

// AddUserAction adds a user.
func AddUserAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	user := &model.User{}
	if err := c.BindJSON(user); nil != err {
		result.Error(err)

		return
	}
	// password
	if 5 <= len(user.Password) {
		user.PasswordHash = crypto.HashAndSalt([]byte(user.Password))
		user.Token = uuid.NewV4().String()
	} else {
		result.Error(errors.New("password required"))

		return
	}
	user.Password = ""

	srv := service.FromContext(c)
	if err := srv.Users.Create(c, user); nil != err {
		result.Error(err)
	}
}

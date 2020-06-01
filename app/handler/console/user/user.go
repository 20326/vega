package user

import (
	"errors"
	"github.com/20326/vega/pkg/crypto"
	uuid "github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"time"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/params"
	"github.com/20326/vega/pkg/render"
	"github.com/gin-gonic/gin"
)

// GetUsersAction gets users.
func GetUsersAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)

	where := ""
	whereArgs := []interface{}{}

	name := c.Query("name")
	if "" != name {
		where += " `username` LIKE ? OR `nickname` LIKE ? "
		whereArgs = append(whereArgs, "%"+name+"%", "%"+name+"%")
	}

	//role := c.Query("role[]")
	//if "" != role {
	//	where += " `role` = ? "
	//	whereArgs = append(whereArgs, role)
	//}

	status := params.GetIntArgs(c, "status")
	if 0 < status {
		where += " `status` = ? "
		whereArgs = append(whereArgs, status)
	}

	pageQuery := model.PageQuery{
		PageNo: params.GetIntArgs(c, "pageNo"),
		PageSize: params.GetIntArgs(c, "pageSize"),
		Where: where,
		WhereArgs: whereArgs,
	}

	users, pagination := srv.Users.FindWhere(pageQuery)

	//var users []*model.ConsoleUser
	//for _, userModel := range userModels {
	//	comment := &model.ConsoleUser{
	//		ID:   userModel.ID,
	//		Name: userModel.Name,
	//	}
	//
	//	users = append(users, comment)
	//}

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

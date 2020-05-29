package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/render"
	"github.com/gin-gonic/gin"
)

// GetUsersAction gets users.
func GetUsersAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	users, _ := srv.Users.List(c)

	//var users []*model.ConsoleUser
	//for _, userModel := range userModels {
	//	comment := &model.ConsoleUser{
	//		ID:   userModel.ID,
	//		Name: userModel.Name,
	//	}
	//
	//	users = append(users, comment)
	//}

	data := map[string]interface{}{}
	data["users"] = users
	result.Result = data
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
	if err := srv.Users.Update(c, user); nil != err {
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

	srv := service.FromContext(c)
	if err := srv.Users.Create(c, user); nil != err {
		result.Error(err)
	}
}

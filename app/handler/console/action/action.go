package action

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/render"
	"github.com/gin-gonic/gin"
)

// GetActionsAction gets actions.
func GetActionsAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	actions, _ := srv.Actions.List(c)

	//var actions []*model.ConsoleAction
	//for _, actionModel := range actionModels {
	//	comment := &model.ConsoleAction{
	//		ID:   actionModel.ID,
	//		Name: actionModel.Name,
	//	}
	//
	//	actions = append(actions, comment)
	//}

	data := map[string]interface{}{}
	data["actions"] = actions
	result.Result = data
}

// GetActionAction get a action.
func GetActionAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	data, err := srv.Actions.Find(c, id)
	if nil == data {
		result.Error(err)

		return
	}

	result.Result = data
}

// DeleteActionAction remove a action.
func DeleteActionAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Actions.Delete(c, id); nil != err {
		result.Error(err)

	}
}

// UpdateActionAction updates a action.
func UpdateActionAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	action := &model.Action{Model: model.Model{ID: uint64(id)}}
	if err := c.BindJSON(action); nil != err {
		result.Error(errors.New("parses update action request failed"))

		return
	}

	srv := service.FromContext(c)
	if err := srv.Actions.Update(c, action); nil != err {
		result.Error(err)
	}
}

// AddActionAction adds a action.
func AddActionAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	action := &model.Action{}
	if err := c.BindJSON(action); nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Actions.Create(c, action); nil != err {
		result.Error(err)
	}
}

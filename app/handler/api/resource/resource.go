package resource

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/render"
	"github.com/gin-gonic/gin"
)

// GetResourcesAction gets resources.
func GetResourcesAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	resources, _ := srv.Resources.List(c)

	//var resources []*model.ConsoleResource
	//for _, resourceModel := range resourceModels {
	//	comment := &model.ConsoleResource{
	//		ID:   resourceModel.ID,
	//		Name: resourceModel.Name,
	//	}
	//
	//	resources = append(resources, comment)
	//}

	data := map[string]interface{}{}
	data["resources"] = resources
	result.Result = data
}

// GetResourceAction get a resource.
func GetResourceAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	data, err := srv.Resources.Find(c, id)
	if nil == data {
		result.Error(err)

		return
	}

	result.Result = data
}

// DeleteResourceAction remove a resource.
func DeleteResourceAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Resources.Delete(c, id); nil != err {
		result.Error(err)

	}
}

// UpdateResourceAction updates a resource.
func UpdateResourceAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	resource := &model.Resource{Model: model.Model{ID: uint64(id)}}
	if err := c.BindJSON(resource); nil != err {
		result.Error(errors.New("parses update resource request failed"))

		return
	}

	srv := service.FromContext(c)
	if err := srv.Resources.Update(c, resource); nil != err {
		result.Error(err)
	}
}

// AddResourceAction adds a resource.
func AddResourceAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	resource := &model.Resource{}
	if err := c.BindJSON(resource); nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Resources.Create(c, resource); nil != err {
		result.Error(err)
	}
}

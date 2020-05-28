package setting

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/render"
	"github.com/gin-gonic/gin"
)

// GetSettingsAction gets settings.
func GetSettingsAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	s := service.FromContext(c)
	settings, _ := s.Settings.List(c)

	//var settings []*model.ConsoleSetting
	//for _, settingModel := range settingModels {
	//	comment := &model.ConsoleSetting{
	//		ID:   settingModel.ID,
	//		Name: settingModel.Name,
	//	}
	//
	//	settings = append(settings, comment)
	//}

	data := map[string]interface{}{}
	data["settings"] = settings
	result.Result = data
}

// GetSettingAction get a setting.
func GetSettingAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	data, err := srv.Settings.Find(c, id)
	if nil == data {
		result.Error(err)

		return
	}

	result.Result = data
}

// DeleteSettingAction remove a setting.
func DeleteSettingAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Settings.Delete(c, id); nil != err {
		result.Error(err)

	}
}

// UpdateSettingAction updates a setting.
func UpdateSettingAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Error(err)

		return
	}

	setting := &model.Setting{Model: model.Model{ID: uint64(id)}}
	if err := c.BindJSON(setting); nil != err {
		result.Error(errors.New("parses update setting request failed"))

		return
	}

	srv := service.FromContext(c)
	if err := srv.Settings.Update(c, setting); nil != err {
		result.Error(err)
	}
}

// AddSettingAction adds a setting.
func AddSettingAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	setting := &model.Setting{}
	if err := c.BindJSON(setting); nil != err {
		result.Error(err)

		return
	}

	srv := service.FromContext(c)
	if err := srv.Settings.Create(c, setting); nil != err {
		result.Error(err)
	}
}

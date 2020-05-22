package setting

import (
	"net/http"
	"strconv"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/errors"
	"github.com/20326/vega/pkg/render"
	// "github.com/phuslu/log"
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
		result.Code = errors.CodeErr

		return
	}

	s := service.FromContext(c)
	data, _ := s.Settings.Find(c, id)
	if nil == data {
		result.Code = errors.CodeErr

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
		result.Code = errors.CodeErr
		result.Msg = err.Error()

		return
	}

	s := service.FromContext(c)
	if err := s.Settings.Delete(c, id); nil != err {
		result.Code = errors.CodeErr
		result.Msg = err.Error()
	}
}

// UpdateSettingAction updates a setting.
func UpdateSettingAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		result.Code = errors.CodeErr
		result.Msg = err.Error()

		return
	}

	setting := &model.Setting{Model: model.Model{ID: uint64(id)}}
	if err := c.BindJSON(setting); nil != err {
		result.Code = errors.CodeErr
		result.Msg = "parses update setting request failed"

		return
	}

	s := service.FromContext(c)
	if err := s.Settings.Update(c, setting); nil != err {
		result.Code = errors.CodeErr
		result.Msg = err.Error()
	}
}

// AddSettingAction adds a setting.
func AddSettingAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	setting := &model.Setting{}
	if err := c.BindJSON(setting); nil != err {
		result.Code = errors.CodeErr
		result.Msg = "parses add setting request failed"

		return
	}

	s := service.FromContext(c)
	if err := s.Settings.Create(c, setting); nil != err {
		result.Code = errors.CodeErr
		result.Msg = err.Error()
	}
}

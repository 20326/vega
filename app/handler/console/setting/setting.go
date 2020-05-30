package setting

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service"
	"github.com/20326/vega/pkg/render"
	"github.com/gin-gonic/gin"
)

// GetSettingsAction gets settings.
func GetSettingsAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	srv := service.FromContext(c)
	settingModels, _ := srv.Settings.List(c)

	settings := map[string]interface{}{}
	for _, settingModel := range settingModels {
		settings[settingModel.Name] = settingModel.Value
	}

	data := map[string]interface{}{}
	data["data"] = settings
	result.Result = data
}

// GetSettingAction get a setting.
func GetSettingAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	idArg := c.Param("id")
	id, err := strconv.ParseUint(idArg, 10, 64)
	if nil != err {
		// ID or Category

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

// UpdateSettingsAction updates a group settings.
func UpdateSettingsAction(c *gin.Context) {
	result := render.NewResult()
	defer c.JSON(http.StatusOK, result)

	args := map[string]interface{}{}
	if err := c.BindJSON(&args); nil != err {
		result.Error(err)
		return
	}

	groupArg := c.Param("group")

	var basics []*model.Setting
	for k, v := range args {
		if !strings.HasPrefix(k, groupArg) {
			continue
		}
		var value interface{}
		switch v.(type) {
		case bool:
			value = strconv.FormatBool(v.(bool))
		case float64:
			value = strconv.FormatFloat(v.(float64), 'f', 0, 64)
		default:
			value = strings.TrimSpace(v.(string))
		}

		basic := &model.Setting{
			Name:  k,
			Value: value.(string),
		}
		basics = append(basics, basic)
	}

	srv := service.FromContext(c)
	if err := srv.Settings.Updates(c, basics); nil != err {
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

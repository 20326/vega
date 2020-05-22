package service

import (
	"github.com/20326/vega/app/config"
	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service/action"
	"github.com/20326/vega/app/service/admission"
	"github.com/20326/vega/app/service/perm"
	"github.com/20326/vega/app/service/resource"
	"github.com/20326/vega/app/service/role"
	"github.com/20326/vega/app/service/setting"
	"github.com/20326/vega/app/service/shared/db"
	"github.com/20326/vega/app/service/user"

	"github.com/gin-gonic/gin"
	"github.com/phuslu/log"
)

const (
	ContextServiceKey = "serviceCtx"
)

type (
	Service struct {
		Actions    model.ActionService
		Admissions model.AdmissionService
		Resources  model.ResourceService
		Perms      model.PermService
		Roles      model.RoleService
		Users      model.UserService
		Settings   model.SettingService
	}
)

func (s *Service) WithContext(c *gin.Context) {
	c.Set(ContextServiceKey, s)
}

func FromContext(c *gin.Context) *Service {
	return c.MustGet(ContextServiceKey).(*Service)
}

func NewService(config *config.Config) *Service {
	// init db
	dbs, err := db.NewDB(db.Config{
		Driver:          config.Database.Driver,
		DSN:             config.Database.DSN,
		TablePrefix:     config.Database.TablePrefix,
		AutoMigrate:     config.Database.AutoMigrate,
		LogMode:         config.Database.LogMode,
		MaxIdleConns:    config.Database.MaxIdleConns,
		MaxOpenConns:    config.Database.MaxOpenConns,
		ConnMaxLifetime: config.Database.ConnMaxLifetime,
	})
	if nil != err {
		log.Fatal().Err(err).Msg("Init db has some errors!")
	}

	// init perm framework
	admissions := admission.New(admission.Config{
		CasbinModel: config.Admission.CasbinModel,
		TablePrefix: config.Admission.TablePrefix,
		LogMode:     config.Admission.LogMode,
	}, dbs)

	// auto migrate

	return &Service{
		Actions:    action.New(dbs),
		Admissions: admissions,
		Resources:  resource.New(dbs),
		Perms:      perm.New(dbs),
		Roles:      role.New(dbs),
		Users:      user.New(dbs),
		Settings:   setting.New(dbs),
	}
}

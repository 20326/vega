package service

import (
	"context"

	"github.com/20326/vega/app/config"
	"github.com/20326/vega/app/model"
	"github.com/20326/vega/app/service/action"
	"github.com/20326/vega/app/service/admission"
	"github.com/20326/vega/app/service/permission"
	"github.com/20326/vega/app/service/resource"
	"github.com/20326/vega/app/service/role"
	"github.com/20326/vega/app/service/setting"
	"github.com/20326/vega/app/service/shared/db"
	"github.com/20326/vega/app/service/user"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	ContextServiceKey = "vega.run.serviceCtx"
)

type (
	Service struct {
		Log         *logrus.Logger
		Actions     model.ActionService
		Admissions  model.AdmissionService
		Resources   model.ResourceService
		Permissions model.PermissionService
		Roles       model.RoleService
		Users       model.UserService
		Settings    model.SettingService
	}
)

func (s *Service) WithContext(c *gin.Context) {
	c.Set(ContextServiceKey, s)
}

func FromContext(c *gin.Context) *Service {
	return c.MustGet(ContextServiceKey).(*Service)
}

func (s *Service) GetLogger() *logrus.Logger{
	return s.Log
}

func NewService(config *config.Config, log *logrus.Logger) *Service {
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
		log.Fatalf("Init db has some errors! error: %s", err)
	}

	// init permission framework
	admissions := admission.New(admission.Config{
		CasbinModel: config.Admission.CasbinModel,
		TablePrefix: config.Admission.TablePrefix,
		LogMode:     config.Admission.LogMode,
	}, dbs)

	// auto migrate

	srv := &Service{
		Log:         log,
		Actions:     action.New(dbs),
		Admissions:  admissions,
		Resources:   resource.New(dbs),
		Permissions: permission.New(dbs),
		Roles:       role.New(dbs),
		Users:       user.New(dbs),
		Settings:    setting.New(dbs),
	}

	var ctx = context.Background()
	roleList, _ := srv.Roles.List(ctx)
	_ = srv.Admissions.LoadAllPolicy(ctx, roleList)

	return srv
}

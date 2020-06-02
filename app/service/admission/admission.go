package admission

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"log"
	"sync"

	"github.com/20326/vega/app/model"
	"github.com/casbin/casbin/v2"
	gormCasbin "github.com/casbin/gorm-adapter/v2"
	"github.com/jinzhu/gorm"
)

// New returns a new AdmissionService.
type admissionService struct {
	db       *gorm.DB
	mutex    *sync.Mutex
	enforcer *casbin.Enforcer
}

var rolePrefix = "role"

func New(config Config, db *gorm.DB) model.AdmissionService {
	var err error
	adapter, err := gormCasbin.NewAdapterByDBUsePrefix(db, config.TablePrefix)
	if nil != err {
		log.Fatalln("new casbin adapter failed")
		return nil
	}

	enforcer, err := casbin.NewEnforcer(config.CasbinModel, adapter)
	enforcer.EnableLog(config.LogMode)
	// log.Fatalln(err)

	return &admissionService{
		db:       db,
		mutex:    &sync.Mutex{},
		enforcer: enforcer,
	}
}

// LoadAllPolicy returns all policy from the datastore.
func (s *admissionService) LoadAllPolicy(ctx context.Context, roles []*model.Role) error {

	logrus.Infof("LoadAllPolicyroles: %+v", roles)
	for _, role := range roles {
		logrus.Infof("LoadAllPolicyroles: s.enforcer: %+v", s.enforcer)
		roleKey := fmt.Sprintf("%s_%d", rolePrefix, role.ID)
		_, _ = s.enforcer.DeleteRole(roleKey)
		// actionIDs := role.GetActionIds()
		for _, action := range role.Actions {
			// actions `[{"action":"add","defaultCheck":false,"describe":"新增"}]`
			for _, resource := range action.Resources {
				_, _ = s.enforcer.AddPermissionForUser(roleKey, resource.Path, resource.Method)
			}
		}
	}

	err := s.enforcer.LoadPolicy()
	if nil != err {
		// log.Fatalln("load casbin policy failed")
		return err
	}
	return nil
}

// DeleteAllPolicy deletes policy from the datastore.
func (s *admissionService) DeleteAllPolicy(ctx context.Context) error {

	return nil
}

func (s *admissionService) Admit(ctx context.Context, user *model.User, subject string, action string) (bool, error) {
	allowed := false
	for _, role := range user.Roles {
		roleKey := fmt.Sprintf("%s_%d", rolePrefix, role.ID)
		allowed, _ = s.enforcer.Enforce(roleKey, subject, action)
		logrus.Warn("Admit for %s %s %s", subject, action, allowed)
		if allowed {
			return true, nil
		}
	}
	// TODO get role permission

	return false, errors.New("not allowed")
}

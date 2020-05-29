package admission

import (
	"context"
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

func New(config Config, db *gorm.DB) model.AdmissionService {
	var err error
	adapter, err := gormCasbin.NewAdapterByDBUsePrefix(db, config.TablePrefix)
	if nil != err {
		log.Fatalln("new casbin adapter failed")
		return nil
	}

	enforcer, err := casbin.NewEnforcer(config.CasbinModel, adapter)
	enforcer.EnableLog(config.LogMode)

	return &admissionService{
		db:       db,
		mutex:    &sync.Mutex{},
		enforcer: enforcer,
	}
}

// LoadAllPolicy returns all policy from the datastore.
func (s *admissionService) LoadAllPolicy(ctx context.Context) error {
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
	username := user.Username
	// TODO get role permission

	return s.enforcer.Enforce(username, subject, action)
}
